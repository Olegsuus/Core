package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/pkg/errors"
)

func (s *UserStorage) StorageGetUser(ctx context.Context, userID string) (models.User, error) {
	var user models.User

	query, args, err := squirrel.
		Select("id", "name", "email", "password", "created_at").
		From("users").
		Where(squirrel.Eq{"id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return user, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса",
		}
	}

	if err = s.db.GetContext(ctx, &user, query, args...); err != nil {
		return user, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить пользователя",
			Status:        400,
		}
	}

	return user, nil
}
