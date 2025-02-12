package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	apperrors "github.com/Olegsuus/Core/pkg/errors"
	"time"
)

func (s *UserStorage) StorageAddUser(ctx context.Context, user *models.User) (string, error) {
	s.l.Info("добавление нового пользователя:", user)
	query, args, err := squirrel.
		Insert("users").
		Columns("name", "email", "password", "created_at").
		Values(user.Name, user.Email, user.Password, time.Now()).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return "", apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса",
			Status:        500,
		}
	}

	var id string
	if err = s.db.GetContext(ctx, &id, query, args...); err != nil {
		return "", apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при регистрации пользователя",
		}
	}

	return id, nil
}
