package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	repo "github.com/gregves/remindme/pkg/repository"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sql.DB
}

func NewRepository(dialect, dsn string, idleConn, maxConn int) (repo.Repository, error) {
	db, err := sql.Open(dialect, dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &repository{db}, nil
}

func (r *repository) Save(reminder *repo.Reminder) error {
	log.Print("IN SAVE")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO reminder (chat_id, chat_message, is_recurrent, is_everyday, weekly_day, monthly_day, annual_date, unique_date, unique_time) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	// log.Print(reminder.UniqueTime)
	// log.Print(reminder.ChatMessage)
	// log.Print(reminder.IsRecurrent)
	// log.Print(reminder.UniqueDate)

	_, err := r.db.ExecContext(
		ctx,
		query,
		reminder.ChatId,
		reminder.ChatMessage,
		reminder.IsRecurrent,
		reminder.IsEveryDay,
		reminder.WeeklyDate,
		reminder.MonthlyDate,
		reminder.AnnualDate,
		reminder.UniqueDate,
		reminder.UniqueTime,
	)
	if err != nil {
		return err
	}
	log.Print(fmt.Sprintf("Reminder for Chat '%d' saved in database", reminder.ChatId))
	return nil
}

func (r *repository) Close() {
	r.db.Close()
}
