package service

import (
	"errors"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gregves/remindme/pkg/constants"
	repo "github.com/gregves/remindme/pkg/repository"
)

// /remindme to check the stock price : October 25 @ 5pm
// /remindme to check the stock price : each Tuesday @ 5pm
// /remindme to check the stock price : each Nov 17 @ 5pm
// /remindme to check the stock price : tomorrow @ 5pm
// /remindme to check the stock price : everyday @ 5pm

type (
	Converter struct {
		Input        string
		TempReminder TempReminder
		Reminder     repo.Reminder
	}
	TempReminder struct {
		Text    string
		DateStr string
		TimeStr string
	}
)

var (
	errMissingPipeSymbol      = errors.New("Missing | symbol to delimitate message from date. Example: '/remindme check the stock price | October 17 @ 2:30PM'")
	errorMissingArobaseSymbol = errors.New("Missing @ symbol to delimitate date from time. Example: '/remindme check the stock price | October 17 @ 5:30AM'")
	erroInvalidDate           = errors.New("Wrong date format or missing. Example. /remindme check the stock price | October 17 (or today or tomorrow or everyday or each Tueday) @ 8:00PM'")
	erroInvalidTime           = errors.New("Wrong time format. Example. /remindme check the stock price | October 17  @ 17:00")
	DatePattern               = "(today|tomorrow)|" +
		"^(each\\s+(Monday|Tuesday|Wednesday|Thirsday|Friday|Saturday|Sunday))|" +
		"^(each\\s+(Jan(uary)?|Feb(ruary)?|Mar(ch)?|Apr(il)?|May|Jun(e)?|Jul(y)?|Aug(ust)?|Sep(tember)?|Oct(ober)?|Nov(ember)?|Dec(ember)?)\\s+\\d{1,2})|" +
		"^(each\\s+(Monday|Tuesday|Wednesday|Thursday|Friday|Saturday|Sunday))|" +
		"^(Jan(uary)?|Feb(ruary)?|Mar(ch)?|Apr(il)?|May|Jun(e)?|Jul(y)?|Aug(ust)?|Sep(tember)?|Oct(ober)?|Nov(ember)?|Dec(ember)?)\\s+\\d{1,2}"
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

func (c *Converter) IsValidInput() bool {
	rawSplit := strings.Split(c.Input, "|")
	if len(rawSplit) != 2 {
		log.Print(errMissingPipeSymbol)
		return false
	}

	splitDateTime := strings.Split(rawSplit[1], "@")
	date := strings.TrimSpace(splitDateTime[0])
	if len(splitDateTime) != 2 {
		log.Print(errorMissingArobaseSymbol)
		return false
	}

	isValidDate, err := patternSearch(DatePattern, date)
	if err != nil {
		log.Print(err)
		return false
	}
	if !isValidDate {
		log.Print(erroInvalidDate)
		return false
	}

	requestedTime := strings.TrimPrefix(splitDateTime[1], " ")
	_, err = time.Parse(time.Kitchen, requestedTime)
	if err != nil {
		log.Print(err)
		return false
	}

	// needed by ToReminder()
	c.TempReminder = TempReminder{
		Text:    rawSplit[0],
		DateStr: date,
		TimeStr: requestedTime,
	}
	return true
}

func patternSearch(pattern string, input string) (bool, error) {
	log.Print(input)
	match, err := regexp.MatchString(pattern, input)
	if err != nil {
		return false, err
	}
	if !match {
		return false, nil
	}
	return true, nil
}

func (c *Converter) ToReminder(chatId int64) error {
	c.Reminder = repo.Reminder{
		ChatId:      chatId,
		ChatMessage: c.TempReminder.Text,
	}
	return nil
}
