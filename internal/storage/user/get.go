package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/pkg/errors"
)

func (s *UserStorage) StorageGetUser(ctx context.Context, userID string) (*models.User, error) {
	query, args, err := squirrel.
		Select("id", "name", "email", "password", "created_at").
		From("users").
		Where(squirrel.Eq{"id": userID}).
		ToSql()

	if err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса",
		}
	}

	var user *models.User
	if err = s.db.SelectContext(ctx, &user, query, args...); err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить пользователя",
			Status:        400,
		}
	}

	return user, nil
}
