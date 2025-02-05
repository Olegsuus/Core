package service

import (
	"context"
	"fmt"
	"log/slog"
)

func (s *PostService) Remove(ctx context.Context, id string) error {
	const op = "service.Remove"

	if err := s.psP.Remove(ctx, id); err != nil {
		s.l.Error("ошибка при удалении поста", slog.String("details", fmt.Sprintf("%s: %w", op, err)))
		return err
	}

	s.l.Info("пост успешно удален")

	return nil
}
