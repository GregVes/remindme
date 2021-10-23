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
	ErrMissingPipeSymbol    = errors.New("Missing | symbol to delimitate message from date. Example: '/remindme check the stock price | October 17 @ 17:00'")
	ErrMissingArobaseSymbol = errors.New("Missing @ symbol to delimitate date from time. Example: '/remindme check the stock price | October 17 @ 08:00'")
	ErrInvalidDate          = errors.New("Wrong date format or missing. Example. /remindme check the stock price | October 17 (or today or tomorrow or everyday or each Tueday) @ 8:00'")
	ErrInvalidTime          = errors.New("Wrong time format. Example. /remindme check the stock price | October 17  @ 17:00")
	ErrInvalidCommand       = errors.New("Not a reminder. Prefix your message with /remindme")
	ErrDb                   = errors.New("An error occurred while trying to save your reminder. Please try later.")

	SuccessSave = "\xE2\x9C\x85 reminder saved"
)
