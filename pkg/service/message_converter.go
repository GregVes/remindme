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
// /remindme to check the stock price : tomorrow @ 5pm
// /remindme to check the stock price : everyday @ 5pm

type (
	Converter struct {
		Input    string
		Reminder repo.Reminder
	}
)

var (
	errMissingPipeSymbol      = errors.New("Missing | symbol to delimitate message from date. Example: '/remindme check the stock price | October 17 @ 5pm'")
	errorMissingArobaseSymbol = errors.New("Missing @ symbol to delimitate date from time. Example: '/remindme check the stock price | October 17 @ 5pm'")
	erroInvalidDate           = errors.New("Wrong date format or missing. Example. /remindme check the stock price | October 17 (or today or tomorrow or everyday or each Tueday) @ 5pm'")
	erroInvalidTime           = errors.New("Wrong time format. Example. /remindme check the stock price | October 17  @ 17:00")
	datePattern               = "(today|tomorrow)|" +
		"each\\s+(Monday|Tuesday|Wednesday|Thirsday|Friday|Saturday|Sunday)|" +
		"(Jan(uary)?|Feb(ruary)?|Mar(ch)?|Apr(il)?|May|Jun(e)?|Jul(y)?|Aug(ust)?|Sep(tember)?|Oct(ober)?|Nov(ember)?|Dec(ember)?)\\s+\\d{1,2}"
	timePattern = "^(0?[1-9]|1[012]):([0-5][0-9])[ap]m$"
)

func IsNewReminder(message string) bool {
	return strings.HasPrefix(message, constants.ReminderCommand)
}

func NewConverter(message string) *Converter {
	return &Converter{
		Input:    message,
		Reminder: repo.Reminder{},
	}
}

func (c *Converter) IsValidInput() bool {
	raw := strings.TrimPrefix(c.Input, constants.ReminderCommand)
	raw = strings.TrimSpace(raw)
	rawSplit := strings.Split(raw, "|")
	if len(rawSplit) != 2 {
		log.Print(errMissingPipeSymbol)
		return false
	}

	splitDateTime := strings.Split(rawSplit[1], "@")
	if len(splitDateTime) != 2 {
		log.Print(errorMissingArobaseSymbol)
		return false
	}

	isValidDate, err := PatternSearch(datePattern, splitDateTime[0])
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

	return true
}

func (c *Converter) ToReminder() error {

	return nil
}

func PatternSearch(pattern string, input string) (bool, error) {
	match, err := regexp.MatchString(pattern, input)
	if err != nil {
		return false, err
	}
	if !match {
		return false, nil
	}
	return true, nil
}
