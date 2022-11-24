package repository

import (
	"context"
	"database/sql"

	"github.com/finnpn/workout-tracker/pkg/log"
	"github.com/finnpn/workout-tracker/usecases/in"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Register(ctx context.Context, record *in.Register) error {

	var (
		insertRecord = "INSERT INTO users (email, hashed_passwd) VALUES (?,?)"
	)

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Error("fail roll back with err =%v", err)
		}
	}()
	_, err = r.db.QueryContext(ctx, insertRecord, record.Email, record.Password)

	if err != nil {
		return err
	}
	return nil

}

func (r *UserRepository) Login(ctx context.Context, record *in.Login) error {
	return nil
}
