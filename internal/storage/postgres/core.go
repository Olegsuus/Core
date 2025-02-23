package storage

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type RepositoryImpl struct {
	db *sqlx.DB
	l  *slog.Logger
}

func NewRepositoryImpl(db *sqlx.DB, l *slog.Logger) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
		l:  l,
	}
}
