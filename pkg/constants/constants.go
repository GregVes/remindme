package constants

import (
	"errors"
)

const (
	MonthDayFormat = "January 2"
	YMDDateFormat  = "2006-01-02"

	TimeFormat      = "15:04"
	ReminderCommand = "/remindme"

	NoUniqueDate = "0001-01-01"
)

var (
	ErrMissingPipeSymbol    = errors.New("\xF0\x9F\x9A\xAB Missing | symbol to delimitate message from date. Example: 'check the stock price | October 17 @ 17:00'")
	ErrMissingArobaseSymbol = errors.New("\xF0\x9F\x86\x98 Missing @ symbol to delimitate date from time. Example: 'check the stock price | each Tuesday @ 08:00'")
	ErrInvalidDate          = errors.New("\xF0\x9F\x9A\xA8 Wrong date format or missing. Example. check the stock price | 2021-10-20 (or each Tuesday or everyday or each October 20 or each 20) @ 8:00'")
	ErrInvalidTime          = errors.New("\xF0\x9F\x9A\xA9 Wrong time format. Minutes should be '00' or '30' Example: 'check the stock price | each October 20  @ 17:00'")
	ErrDb                   = errors.New("\xF0\x9F\x9A\xA7 An error occurred while trying to save your reminder. Please try later or contact Greg at +491749505953")
	ErrNotAReminder         = errors.New("\xE2\x9D\x97 This is not a valid reminder. Prefix your message with /remindme")
	ErrDbDuplicate          = errors.New("\xF0\x9F\x98\xB3 You already asked me to remind you this. So ask me something else.")

	SuccessSave = "\xE2\x9C\x85 I saved the info and will remind you about it when it is time to!"
)
