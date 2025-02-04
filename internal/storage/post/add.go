package storage

import (
	"context"
	"github.com/Olegsuus/Core/internal/domain/post"
	apperrors "github.com/Olegsuus/Core/internal/errors"
	"log"
	"time"
)

func (s *PostStorage) Add(ctx context.Context, post *domain.Post) (int64, error) {
	const op = "storage.Add"

	query := `INSERT INTO posts (title, content, created_at) VALUES ($1, $2, $3) RETURNING id`

	var id int64
	err := s.pg.QueryRow(ctx, query, post.Title, post.Content, time.Now()).Scan(&id)
	if err != nil {
		log.Printf("ошибка при сохранении нового поста: (%s: %w)", op, err)
		return 0, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при сохранении нового поста",
		}
	}

	log.Print("Новый пост успешно добавлен")

	return id, nil
}
