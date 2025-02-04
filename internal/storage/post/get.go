package storage

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/domain/dto"
	"github.com/Olegsuus/Core/internal/domain/post"
	"github.com/Olegsuus/Core/internal/errors"
	"log"
)

func (s *PostStorage) GetMany(ctx context.Context, settings dto.GetManyPostSettings) ([]*domain.Post, error) {
	const op = "storage.GetMany"

	order := "ASC"
	if settings.SortDesc {
		order = "DESC"
	}

	query := fmt.Sprintf(
		"SELECT id, title, content, created_at FROM posts ORDER BY created_at %s LIMIT $1 OFFSET $2",
		order,
	)

	rows, err := s.pg.Query(ctx, query, settings.Limit, settings.Offset)
	if err != nil {
		log.Printf("ошибка при получении списка постов: (%s: %w)", op, err)
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить список постов",
			Status:        400,
		}
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		var post domain.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			log.Printf("ошибка при декодировании поста: (%s: %w)", op, err)
			return nil, errors.AppError{
				BusinessError: err.Error(),
				UserError:     "ошибка при получении списка постов",
				Status:        500,
			}
		}
		posts = append(posts, &post)
	}

	log.Print("список постов успешно получен")

	return posts, nil
}
