package postgresql

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	repo "github.com/gregves/remindme/pkg/repository"
	"github.com/stretchr/testify/assert"
)

var reminder = &repo.Reminder{
	ChatId:                      1111,
	ChatMessage:                 "super message",
	IsRecurrent:                 true,
	IsEveryDay:                  true,
	RecurrentWeekDay:            "Monday",
	RecurrentMonthlyDatePattern: nil,
	RecurrentAnnualDate:         "October 17",
	UniqueDate:                  "",
	UniqueTime:                  nil,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestSaveReminderInDB(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := `INSERT INTO reminder (chat_id, chat_message, is_recurrent, is_everyday, recurrent_week_day, recurrent_month_day, recurrent_date, unique_date, unique_time) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(reminder.ChatId, reminder.ChatMessage, reminder.IsRecurrent, reminder.IsEveryDay, reminder.RecurrentWeekDay, reminder.RecurrentMonthlyDatePattern, reminder.RecurrentAnnualDate, reminder.UniqueDate, reminder.UniqueTime).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Save(reminder)
	assert.NoError(t, err)
}
