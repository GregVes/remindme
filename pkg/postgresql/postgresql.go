package postgresql

import (
	"context"
	"database/sql"
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

	query := `INSERT INTO reminder (chat_id, text, recurrent_day, target_date, target_time) VALUES(?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, reminder.ChatId, reminder.Text, reminder.RecurrentDay, reminder.TargetDate, reminder.TargetTime)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Close() {
	r.db.Close()
}
