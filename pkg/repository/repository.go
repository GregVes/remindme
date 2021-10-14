package repository

import (
	"time"

	_ "github.com/lib/pq"
)

type (
	Reminder struct {
		Id           string
		ChatId       int
		Text         string
		RecurrentDay string
		TargetDate   time.Time
		TargetTime   time.Time
	}

	Repository interface {
		Save(reminder *Reminder) error
	}
)
