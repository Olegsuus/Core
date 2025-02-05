package service

import (
	"context"
	"github.com/Olegsuus/Core/internal/domain/dto"
	"github.com/Olegsuus/Core/internal/domain/post"
	"log/slog"
)

type PostService struct {
	psP postStorageProvider
	l   *slog.Logger
}

type postStorageProvider interface {
	Add(ctx context.Context, post *domain.Post) (string, error)
	GetMany(ctx context.Context, settings dto.GetManyPostSettings) ([]*domain.Post, error)
	Remove(ctx context.Context, id string) error
}

func RegisterPostService(psP postStorageProvider, l *slog.Logger) *PostService {
	return &PostService{
		psP: psP,
		l:   l,
	}
}
