package postgresql

import (
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	repo "github.com/gregves/remindme/pkg/repository"
	"github.com/stretchr/testify/assert"
)

var reminder = &repo.Reminder{
	Id:           uuid.New().String(),
	ChatId:       1111111,
	Text:         "this is a reminder",
	RecurrentDay: "Monday",
	TargetDate:   time.Now(),
	TargetTime:   time.Now(),
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()

	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestSaveReminder(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := `INSERT INTO reminder (chat_id, text, recurrent_day, target_date, target_time) VALUES(?, ?, ?, ?, ?)`

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(reminder.ChatId, reminder.Text, reminder.RecurrentDay, reminder.TargetDate, reminder.TargetTime).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Save(reminder)
	assert.NoError(t, err)
}
