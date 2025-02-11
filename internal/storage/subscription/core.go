package storage

import (
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type SubscriptionStorage struct {
	db *sqlx.DB
	l  *slog.Logger
}

func RegisterNewSubscriptionStorage(db *sqlx.DB, l *slog.Logger) *SubscriptionStorage {
	return &SubscriptionStorage{
		db: db,
		l:  l,
	}
}
