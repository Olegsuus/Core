package storage

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/storage"
	"github.com/Olegsuus/Core/pkg/errors"
	"log/slog"
)

func (s *SubscriptionStorage) Unsubscribe(ctx context.Context, subscriptionEntity *storage.SubscriptionEntity) error {
	s.l.Info("Создание новой подписки", slog.String("details",
		fmt.Sprintf("пользователь: %s, отписывается от %s", subscriptionEntity.SubscriberID,
			subscriptionEntity.SubscribedToID)))

	query, args, err := squirrel.
		Delete("subscriptions").
		Where(squirrel.Eq{"subscriber_id": subscriptionEntity.SubscriberID,
			"subscribed_to_id": subscriptionEntity.SubscribedToID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса на отписку",
		}
	}

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при отписке",
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось определить результат отписки",
		}
	}
	if rowsAffected == 0 {
		return errors.AppError{
			BusinessError: fmt.Sprintf("не найдена подписка для subscriber_id=%s и subscribed_to_id=%s",
				subscriptionEntity.SubscriberID, subscriptionEntity.SubscribedToID),
			UserError: "подписка не найдена",
			Status:    404,
		}
	}

	return nil
}
