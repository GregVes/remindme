package repository

import (
	_ "github.com/lib/pq"
)

type (
	Reminder struct {
		Id          string `json:"id"`
		ChatId      int    `json:"chat_id"`
		ChatMessage string `json:"chat_message"`
		IsRecurrent bool   `json:"is_recurrent"`
		IsEveryDay  bool   `json:"is_everyday"`
		WeeklyDate  string `json:"weekly_day"`
		MonthlyDate *int64 `json:"monthly_day"`
		AnnualDate  string `json:"annual_date"`
		UniqueDate  string `json:"unique_date"`
		UniqueTime  string `json:"unique_time"`
	}
	Reminders struct {
		Data []Reminder `json:"data"`
	}

	Repository interface {
		Save(reminder *Reminder) error
		Close()
	}
)
