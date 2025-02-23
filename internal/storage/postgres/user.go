package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	apperrors "github.com/Olegsuus/Core/pkg/errors"
)

func (s *RepositoryImpl) AddUser(ctx context.Context, userParam AddUserParam) (*UserEntity, error) {
	s.l.Info("добавление нового пользователя:", userParam)

	userEntity := UserEntity{
		Name:     userParam.Name,
		Email:    userParam.Email,
		Password: userParam.Password,
	}

	query, args := squirrel.
		Insert("users").
		Columns("name", "email", "password").
		Values(userEntity.Name, userEntity.Email, userEntity.Password).
		Suffix("RETURNING id, name, email, password").
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	if err := s.db.GetContext(ctx, &userEntity, query, args...); err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при регистрации пользователя",
		}
	}

	return &userEntity, nil
}

func (s *RepositoryImpl) GetUser(ctx context.Context, userID string) (*UserEntity, error) {
	query, args := squirrel.
		Select("id", "name", "email", "password", "created_at").
		From("users").
		Where(squirrel.Eq{"id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	var userEntity UserEntity
	if err := s.db.GetContext(ctx, &userEntity, query, args...); err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить пользователя",
			Status:        404,
		}
	}

	return &userEntity, nil
}
