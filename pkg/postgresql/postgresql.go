package postgresql

import (
	"context"
	"database/sql"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `INSERT INTO reminder (chat_id, chat_message, is_recurrent, is_everyday, recurrent_week_day, recurrent_month_day, recurrent_date, unique_date, unique_time) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(
		ctx,
		query,
		reminder.ChatId,
		reminder.ChatMessage,
		reminder.IsRecurrent,
		reminder.IsEveryDay,
		reminder.RecurrentWeekDay,
		reminder.RecurrentMonthlyDatePattern,
		reminder.RecurrentAnnualDate,
		reminder.UniqueDate,
		reminder.UniqueTime,
	)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func (r *repository) Close() {
	r.db.Close()
}
