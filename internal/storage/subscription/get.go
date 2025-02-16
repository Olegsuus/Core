package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/pkg/errors"
)

func (s *SubscriptionStorage) StorageGetSubscribers(ctx context.Context, userID string, settings models.GetManySettings,
) ([]models.User, error) {

	query, args, err := squirrel.
		Select("u.id", "u.name", "u.email", "u.created_at").
		From("subscriptions s").
		Join("users u ON s.subscriber_id = u.id").
		Where(squirrel.Eq{"s.subscribed_to_id": userID}).
		OrderBy("s.created_at DESC").
		Limit(uint64(settings.Limit)).
		Offset(uint64(settings.Offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса на получение подписчиков",
		}
	}

	var subscribers []models.User
	if err = s.db.SelectContext(ctx, &subscribers, query, args...); err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить подписчиков",
			Status:        400,
		}
	}

	return subscribers, nil
}
