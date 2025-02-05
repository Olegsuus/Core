package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/domain/post"
	"log/slog"
)

func (s *PostService) Add(ctx context.Context, title, content string) (string, error) {
	const op = "service.Add"

	newPost := &domain.Post{
		Title:   title,
		Content: content,
	}

	id, err := s.psP.Add(ctx, newPost)
	if err != nil {
		s.l.Error("ошибка при добавлении нового поста", slog.String("details", fmt.Sprintf("%s: %w", op, err)))
		return "", err
	}

	s.l.Info("пост успешно получен")

	return id, nil
}
