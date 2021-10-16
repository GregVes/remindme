package postgresql

import (
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	repo "github.com/gregves/remindme/pkg/repository"
	"github.com/stretchr/testify/assert"
)

var reminder = &repo.Reminder{
	ChatId:              1111,
	ChatMessage:         "super message",
	IsRecurrent:         true,
	TargetRecurrentDate: time.Now(),
	TargetRecurrentDay:  "Monday",
	TargetDate:          time.Now(),
	TargetTime:          "1:30",
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

	query := `INSERT INTO reminder (chat_id, chat_message, is_recurrent, target_recurrent_date, target_recurrent_day, target_date, target_time) VALUES(?, ?, ?, ?, ?, ?, ?)`

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(reminder.ChatId, reminder.ChatMessage, reminder.IsRecurrent, reminder.TargetRecurrentDate, reminder.TargetRecurrentDay, reminder.TargetDate, reminder.TargetTime).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Save(reminder)
	assert.NoError(t, err)
}
