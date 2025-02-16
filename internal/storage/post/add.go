package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	apperrors "github.com/Olegsuus/Core/pkg/errors"
)

func (s *PostStorage) AddPost(ctx context.Context, postEntity storage.PostEntity) (*models.Post, error) {
	s.l.Info("создание нового поста:", postEntity)

	query, args, err := squirrel.
		Insert("posts").
		Columns("user_id", "title", "content").
		Values(postEntity.UserID, postEntity.Title, postEntity.Content).
		Suffix("RETURNING id, user_id, title, content, created_at").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при добавлении поста",
		}
	}

	if err := s.db.GetContext(ctx, &postEntity, query, args...); err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при сохранении нового поста",
		}
	}

	return postEntityToModels(postEntity), nil
}
