package repository

import (
	"time"

	_ "github.com/lib/pq"
)

type (
	Reminder struct {
		Id                  string
		ChatId              int64
		ChatMessage         string
		IsRecurrent         bool
		TargetRecurrentDate time.Time
		TargetRecurrentDay  string
		TargetDate          time.Time
		TargetTime          string
	}

	Repository interface {
		Save(reminder *Reminder) error
	}
)
