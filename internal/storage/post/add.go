package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	apperrors "github.com/Olegsuus/Core/pkg/errors"
	"time"
)

func (s *PostStorage) StorageAddPost(ctx context.Context, post *models.Post) (string, error) {
	s.l.Info("создание нового поста:", post)

	query, args, err := squirrel.
		Insert("posts").
		Columns("user_id", "title", "content", "created_at").
		Values(post.UserID, post.Title, post.Content, time.Now()).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return "", apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при добавлении поста",
		}
	}

	var id string
	if err := s.db.GetContext(ctx, &id, query, args...); err != nil {
		return "", apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при сохранении нового поста",
		}
	}

	return id, nil
}
