package storage

import (
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type UserStorage struct {
	db *sqlx.DB
	l  *slog.Logger
}

func NewUserStorage(db *sqlx.DB, l *slog.Logger) *UserStorage {
	return &UserStorage{
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
