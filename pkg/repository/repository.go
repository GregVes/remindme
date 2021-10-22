package repository

import (
	"time"

	_ "github.com/lib/pq"
)

type (
	Reminder struct {
		Id                          string
		ChatId                      int
		ChatMessage                 string
		IsRecurrent                 bool
		IsEveryDay                  bool
		RecurrentWeekDay            string
		RecurrentMonthlyDatePattern *int64
		RecurrentAnnualDate         string
		UniqueDate                  string
		UniqueTime                  *time.Time
	}

	Repository interface {
		Save(reminder *Reminder) error
	}
)
