package service

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gregves/remindme/pkg/constants"
	repo "github.com/gregves/remindme/pkg/repository"
)

type (
	Converter struct {
		Input        string
		TempReminder TempReminder
		Reminder     repo.Reminder
	}
	TempReminder struct {
		Text        string
		DateStr     string
		TimeStr     string
		IsRecurrent bool
		IsEveryDay  bool
	}
)

var (
	DatePatterns = "(today|everyday)|" +
		"^(each\\s+(Jan(uary)?|Feb(ruary)?|Mar(ch)?|Apr(il)?|May|Jun(e)?|Jul(y)?|Aug(ust)?|Sep(tember)?|Oct(ober)?|Nov(ember)?|Dec(ember)?)\\s+\\d{1,2})|" +
		"^(each\\s+(Monday|Tuesday|Wednesday|Thursday|Friday|Saturday|Sunday))|" +
		"^(Jan(uary)?|Feb(ruary)?|Mar(ch)?|Apr(il)?|May|Jun(e)?|Jul(y)?|Aug(ust)?|Sep(tember)?|Oct(ober)?|Nov(ember)?|Dec(ember)?)\\s+\\d{1,2}|" +
		"^(each\\s+[0-30]{1,2}$)|" +
		`\d{4}-\d{2}-\d{2}`
	WeeklyDatePattern  = "^(each\\s+(Monday|Tuesday|Wednesday|Thursday|Friday|Saturday|Sunday))"
	MonthlyDatePattern = "^(each\\s+[0-30])"

	TimePattern = "^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$"
)

func NewConverter(message string) *Converter {
	return &Converter{
		Input:    message,
		Reminder: repo.Reminder{},
	}
}

func IsNewReminder(message string) bool {
	return strings.HasPrefix(message, constants.ReminderCommand)
}

func (c *Converter) ExtractRawReminder() {
	raw := strings.TrimPrefix(c.Input, constants.ReminderCommand)
	raw = strings.TrimSpace(raw)
	c.Input = raw
}

func (c *Converter) IsValidInput() error {
	// no pipe separating reminder text from date + time
	rawSplit := strings.Split(c.Input, "|")
	if len(rawSplit) != 2 {
		log.Print(constants.ErrMissingPipeSymbol)
		return constants.ErrMissingPipeSymbol
	}
	// missing @ separator between date and time
	splitDateTime := strings.Split(rawSplit[1], "@")
	date := strings.TrimSpace(splitDateTime[0])
	if len(splitDateTime) != 2 {
		log.Print(constants.ErrMissingArobaseSymbol)
		return constants.ErrMissingArobaseSymbol
	}

	// invalid date format
	isValidDate := patternSearch(DatePatterns, date)
	if !isValidDate {
		log.Print(constants.ErrInvalidDate)
		return constants.ErrInvalidDate
	}

	// invalid time format
	time := splitDateTime[1]
	time = strings.TrimSpace(time)
	isValidTime := patternSearch(TimePattern, time)
	if !isValidTime {
		log.Print(constants.ErrInvalidTime)
		return constants.ErrInvalidTime
	}

	// needed by ToReminder()
	c.TempReminder = TempReminder{
		Text:        strings.TrimSpace(rawSplit[0]),
		DateStr:     date,
		TimeStr:     time,
		IsRecurrent: strings.Contains(date, "each") || strings.Contains(date, "everyday"),
		IsEveryDay:  strings.Contains(date, "everyday"),
	}

	return nil
}

func patternSearch(pattern string, input string) bool {
	regExp := regexp.MustCompile(pattern)
	match := regExp.MatchString(input)

	return match
}

func ToValidAnnualDate(dateStr string) string {
	return strings.Replace(dateStr, "each ", "", 1)
}

func ToValidDate(layout string, timeStr string) *time.Time {
	var res time.Time
	res, _ = time.Parse(layout, timeStr)
	return &res
}

func (c *Converter) ToReminder(chatId int) error {
	c.Reminder.ChatId = chatId
	c.Reminder.ChatMessage = c.TempReminder.Text
	c.Reminder.UniqueTime = c.TempReminder.TimeStr

	isToday := strings.Contains(c.TempReminder.DateStr, "today")
	if isToday {
		c.Reminder.UniqueDate = time.Now().Format(constants.YMDDateFormat)
		return nil
	}

	isUniqueDate := !c.TempReminder.IsRecurrent
	if isUniqueDate {
		c.Reminder.UniqueDate = c.TempReminder.DateStr
		return nil
	}

	c.Reminder.UniqueDate = constants.NoUniqueDate
	c.Reminder.IsRecurrent = true

	isDaily := c.TempReminder.IsEveryDay
	isWeekly := patternSearch(WeeklyDatePattern, c.TempReminder.DateStr)
	isMonthly := patternSearch(MonthlyDatePattern, c.TempReminder.DateStr)

	if isDaily {
		c.Reminder.IsEveryDay = true
	} else if isWeekly {
		c.Reminder.WeeklyDate = strings.Replace(c.TempReminder.DateStr, "each ", "", 1)
	} else if isMonthly {
		var monthDay int64
		monthDay, _ = strconv.ParseInt(ToValidAnnualDate(c.TempReminder.DateStr), 0, 0)
		c.Reminder.MonthlyDate = &monthDay
		// annual
	} else {
		c.Reminder.AnnualDate = strings.Replace(c.TempReminder.DateStr, "each ", "", 1)
	}

	return nil
}
