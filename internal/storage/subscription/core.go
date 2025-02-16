package storage

import (
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type SubscriptionStorage struct {
	db *sqlx.DB
	l  *slog.Logger
}

func NewSubscriptionStorage(db *sqlx.DB, l *slog.Logger) *SubscriptionStorage {
	return &SubscriptionStorage{
		db: db,
		l:  l,
	}
}

func userEntityToModels(entity storage.UserEntity) *models.User {
	return &models.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
	}
}
