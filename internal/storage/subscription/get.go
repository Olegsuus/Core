package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	"github.com/Olegsuus/Core/pkg/errors"
	"github.com/Olegsuus/Core/pkg/utils"
)

func (s *SubscriptionStorage) GetSubscribers(ctx context.Context, userID string, limit, offset int,
) ([]*models.User, error) {

	query, args, err := squirrel.
		Select("u.id", "u.name", "u.email", "u.created_at").
		From("subscriptions s").
		Join("users u ON s.subscriber_id = u.id").
		Where(squirrel.Eq{"s.subscribed_to_id": userID}).
		OrderBy("s.created_at DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса на получение подписчиков",
		}
	}

	var usersEntity []storage.UserEntity
	if err = s.db.SelectContext(ctx, &usersEntity, query, args...); err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить подписчиков",
			Status:        404,
		}
	}

	users := utils.MapAsync(usersEntity, userEntityToModels)

	return users, err
}

func userEntityToModels(entity storage.UserEntity) *models.User {
	return &models.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
	}
}
