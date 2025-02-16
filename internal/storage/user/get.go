package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	"github.com/Olegsuus/Core/pkg/errors"
)

func (s *UserStorage) GetUser(ctx context.Context, userID string) (*models.User, error) {
	query, args, err := squirrel.
		Select("id", "name", "email", "password", "created_at").
		From("users").
		Where(squirrel.Eq{"id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса",
		}
	}

	var userEntity storage.UserEntity
	if err = s.db.GetContext(ctx, &userEntity, query, args...); err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить пользователя",
			Status:        404,
		}
	}

	return userEntityToModels(userEntity), nil
}
