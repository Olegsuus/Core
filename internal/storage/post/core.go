package storage

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type PostStorage struct {
	db *sqlx.DB
	l  *slog.Logger
}

func RegisterNewPostStorage(db *sqlx.DB, l *slog.Logger) *PostStorage {
	return &PostStorage{
		db: db,
		l:  l,
	}
}
