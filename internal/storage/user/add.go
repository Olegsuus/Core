package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	apperrors "github.com/Olegsuus/Core/pkg/errors"
)

func (s *UserStorage) AddUser(ctx context.Context, userEntity storage.UserEntity) (*models.User, error) {
	s.l.Info("добавление нового пользователя:", userEntity)
	query, args, err := squirrel.
		Insert("users").
		Columns("name", "email", "password").
		Values(userEntity.Name, userEntity.Email, userEntity.Password).
		Suffix("RETURNING id, name, email, password").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса",
			Status:        500,
		}
	}

	if err = s.db.GetContext(ctx, &userEntity, query, args...); err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при регистрации пользователя",
		}
	}

	return userEntityToModels(userEntity), nil
}
