package storage

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type UserStorage struct {
	db *sqlx.DB
	l  *slog.Logger
}

func RegisterNewUserStorage(db *sqlx.DB, l *slog.Logger) *UserStorage {
	return &UserStorage{
		db: db,
		l:  l,
	}
}
