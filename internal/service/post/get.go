package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/domain/dto"
	domain "github.com/Olegsuus/Core/internal/domain/post"
	"log/slog"
)

func (s *PostService) GetMany(ctx context.Context, settings dto.GetManyPostSettings) ([]*domain.Post, error) {
	const op = "service.GetMany"

	if settings.Limit <= 0 {
		settings.Limit = 15
	}

	if settings.Offset <= 0 {
		settings.Offset = 1
	}

	posts, err := s.psP.GetMany(ctx, settings)
	if err != nil {
		s.l.Error("ошибка при получении списка постов", slog.String("details", fmt.Sprintf("%s: %w", op, err)))
		return nil, err
	}

	s.l.Info("список постов успешно получен")

	return posts, nil
}
