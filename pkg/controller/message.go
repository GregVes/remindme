package controller

import (
	"net/http"
	"os"
)

var TELEGRAM_BOT_TOKEN = os.Getenv("TELEGRAM_BOT_TOKEN")

type (
	Reminder struct {
	}
	storage interface {
		SaveReminder(*Reminder)
	}
)

func PostMessage(w http.ResponseWriter, r *http.Request) {
	// map request to Update
	// get text
	// parse text into a Reminder object (type + message + occurence + time)
	// store reminder into db
}
