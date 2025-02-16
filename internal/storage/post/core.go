package storage

import (
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type PostStorage struct {
	db *sqlx.DB
	l  *slog.Logger
}

func NewPostStorage(db *sqlx.DB, l *slog.Logger) *PostStorage {
	return &PostStorage{
		db: db,
		l:  l,
	}
}

func postEntityToModels(entity storage.PostEntity) *models.Post {
	return &models.Post{
		ID:        entity.ID,
		UserID:    entity.UserID,
		Content:   entity.Content,
		Title:     entity.Title,
		CreatedAt: entity.CreatedAt,
	}
}
