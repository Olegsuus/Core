package storage

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/storage"
	apperrors "github.com/Olegsuus/Core/pkg/errors"
	"log/slog"
)

func (s *SubscriptionStorage) Subscribe(ctx context.Context, subscriptionEntity *storage.SubscriptionEntity) error {
	s.l.Info("Создание новой подписки", slog.String("details",
		fmt.Sprintf("пользователь: %s, подписывается на %s",
			subscriptionEntity.SubscriberID, subscriptionEntity.SubscribedToID)))

	query, args, err := squirrel.
		Insert("subscriptions").
		Columns("subscriber_id", "subscribed_to_id").
		Values(subscriptionEntity.SubscriberID, subscriptionEntity.SubscribedToID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса на подписку",
			Status:        500,
		}
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при подписке",
		}
	}

	return nil
}
