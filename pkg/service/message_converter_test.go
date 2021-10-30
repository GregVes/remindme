package service

import (
	"testing"
	"time"

	"github.com/gregves/remindme/pkg/constants"
	repo "github.com/gregves/remindme/pkg/repository"
	"github.com/stretchr/testify/assert"
)

// func TestIsNewReminder(t *testing.T) {

// 	tests := []struct {
// 		input string
// 		want  bool
// 	}{
// 		{
// 			input: "check the stock price | October 17 @ 13:00",
// 			want:  true,
// 		},
// 		{
// 			input: "check the stock price | October 17 @ 08:00",
// 			want:  false,
// 		},
// 		{
// 			input: "check the stock price | October 17 @ 15:30 /remindme",
// 			want:  false,
// 		},
// 	}
// 	for _, tc := range tests {
// 		got := IsNewReminder(tc.input)
// 		assert.Equal(t, tc.want, got)
// 	}
// }
func TestIsValidInput(t *testing.T) {
	tests := []struct {
		input string
		want  error
	}{
		{
			input: "check the stock price | October 26 @ 23:00",
			want:  nil,
		},
		{
			input: "check the stock price | 2021-10-31 @ 23:00",
			want:  nil,
		},
		{
			input: "check the stock price | October 17 @ 23:00",
			want:  nil,
		},
		{
			input: "check the stock price | each October 17 @ 23:00",
			want:  nil,
		},
		{
			input: "check the stock price | each Tuesday @ 23:00",
			want:  nil,
		},
		{
			input: "check the stock price | each uesday @ 23:00",
			want:  constants.ErrInvalidDate,
		},
		{
			input: "check the stock price October 17 @ 23:00",
			want:  constants.ErrMissingPipeSymbol,
		},
		{
			input: "check the stock price | October 17 23:00",
			want:  constants.ErrMissingArobaseSymbol,
		},
		{
			input: "October 18 | check the stock price @ 23:00",
			want:  constants.ErrInvalidDate,
		},
		{
			input: "check the stock price October | 17 @ 23:00",
			want:  constants.ErrInvalidDate,
		},
		{
			input: "check the stock price | October 18 @ :00",
			want:  constants.ErrInvalidTime,
		},
		{
			input: "@ 15:00 check the stock price",
			want:  constants.ErrMissingPipeSymbol,
		},
	}
	for _, tc := range tests {
		converter := NewConverter(tc.input)
		got := converter.IsValidInput()
		assert.Equal(t, tc.want, got, tc.input)
	}
}

func TestPatternSearch(t *testing.T) {
	tests := []struct {
		pattern string
		input   string
		want    bool
	}{
		{
			pattern: "(hello)",
			input:   "hello",
			want:    true,
		},
		{
			pattern: "(each\\s+sunday)",
			input:   "each sunday",
			want:    true,
		},
		{
			pattern: "(each\\s+sunday)",
			input:   "eachsunday",
			want:    false,
		},
		{
			pattern: "October 26",
			input:   "26 October",
			want:    false,
		},
		{
			pattern: "(tomorrow)|each\\s+sunday",
			input:   "tomorrow",
			want:    true,
		},
		{
			pattern: DatePatterns,
			input:   "each Nov 19",
			want:    true,
		},
		{
			pattern: DatePatterns,
			input:   "each Nov 10",
			want:    true,
		},
		{
			pattern: DatePatterns,
			input:   "eacc Nov 10",
			want:    false,
		},
		{
			pattern: DatePatterns,
			input:   "each Nov10",
			want:    false,
		},
		{
			pattern: DatePatterns,
			input:   "each 10",
			want:    true,
		},
		{
			pattern: DatePatterns,
			input:   "each 10 Nov",
			want:    false,
		},
	}
	for _, tc := range tests {
		got := patternSearch(tc.pattern, tc.input)
		assert.Equal(t, tc.want, got, tc.input)
	}
}

// func TestExtractRawReminder(t *testing.T) {
// 	tests := []struct {
// 		input string
// 		want  string
// 	}{
// 		{
// 			input: "check the stock price | October 17 @ 23:00",
// 			want:  "check the stock price | October 17 @ 23:00",
// 		},
// 		{
// 			input: "buy bread | tomorrow @ 16:30",
// 			want:  "buy bread | tomorrow @ 16:30",
// 		},
// 	}
// 	for _, tc := range tests {
// 		converter := NewConverter(tc.input)
// 		converter.ExtractRawReminder()
// 		assert.Equal(t, tc.want, converter.Input)
// 	}
// }
func TestToValidAnnualDate(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "each October 17",
			want:  "October 17",
		},
		{
			input: "each 17",
			want:  "17",
		},
	}
	for _, tc := range tests {
		got := ToValidAnnualDate(tc.input)
		assert.Equal(t, tc.want, got)
	}
}
func TestToReminder(t *testing.T) {
	var recurrenMonthdate int64 = 20
	uniqueTime := "12:00"
	// scenarii:
	// to check the stock price : everyday @ 5pm OK
	// to check the stock price : each 20 @ 5pm

	// to check the stock price : each Tuesday @ 5pm OK
	// to check the stock price : each November 20 @ 5pm OK
	// to check the stock price : 2021-10-25 @ 5pm OK

	tests := []struct {
		input string
		want  repo.Reminder
	}{
		{
			input: "check the stock price | each October 17 @ 12:00",
			want: repo.Reminder{
				ChatId:      1111111,
				ChatMessage: "check the stock price",
				IsRecurrent: true,
				IsEveryDay:  false,
				WeeklyDate:  "",
				MonthlyDate: nil,
				AnnualDate:  "October 17",
				UniqueDate:  constants.NoUniqueDate,
				UniqueTime:  uniqueTime,
			},
		},
		{
			input: "check the stock price | each 20 @ 12:00",
			want: repo.Reminder{
				ChatId:      1111111,
				ChatMessage: "check the stock price",
				IsRecurrent: true,
				IsEveryDay:  false,
				WeeklyDate:  "",
				MonthlyDate: &recurrenMonthdate,
				AnnualDate:  "",
				UniqueDate:  constants.NoUniqueDate,
				UniqueTime:  uniqueTime,
			},
		},
		{
			input: "check the stock price | 2021-10-20 @ 12:00",
			want: repo.Reminder{
				ChatId:      1111111,
				ChatMessage: "check the stock price",
				IsRecurrent: false,
				UniqueDate:  "2021-10-20",
				UniqueTime:  uniqueTime,
			},
		},
		{
			input: "check the stock price | each Wednesday @ 12:00",
			want: repo.Reminder{
				ChatId:      1111111,
				ChatMessage: "check the stock price",
				IsRecurrent: true,
				IsEveryDay:  false,
				WeeklyDate:  "Wednesday",
				MonthlyDate: nil,
				AnnualDate:  "",
				UniqueDate:  constants.NoUniqueDate,
				UniqueTime:  uniqueTime,
			},
		},
		{
			input: "check the stock price | everyday @ 12:00",
			want: repo.Reminder{
				ChatId:      1111111,
				ChatMessage: "check the stock price",
				IsRecurrent: true,
				IsEveryDay:  true,
				WeeklyDate:  "",
				MonthlyDate: nil,
				AnnualDate:  "",
				UniqueDate:  constants.NoUniqueDate,
				UniqueTime:  uniqueTime,
			},
		},
		{
			input: "check the stock price | today @ 12:00",
			want: repo.Reminder{
				ChatId:      1111111,
				ChatMessage: "check the stock price",
				IsRecurrent: false,
				IsEveryDay:  false,
				WeeklyDate:  "",
				MonthlyDate: nil,
				AnnualDate:  "",
				UniqueDate:  time.Now().Format(constants.YMDDateFormat),
				UniqueTime:  uniqueTime,
			},
		},
	}
	for _, tc := range tests {
		converter := NewConverter(tc.input)
		converter.IsValidInput()
		got := converter.ToReminder(tc.want.ChatId)
		if got != nil {
			t.Fatal(got)
		}
		assert.Equal(t, tc.want, converter.Reminder, tc.input)
	}
}
