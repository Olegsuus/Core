package storage

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperrors "github.com/Olegsuus/Core/pkg/errors"
	"log/slog"
	"time"
)

func (s *SubscriptionStorage) StorageAddSubscribe(ctx context.Context, userID, subscriberToId string) error {
	s.l.Info("Создание новой подписки", slog.String("details",
		fmt.Sprintf("пользователь: %s, подписывается на %s", userID, subscriberToId)))

	query, args, err := squirrel.
		Insert("subscriptions").
		Columns("subscriber_id", "subscribed_to_id", "created_at").
		Values(userID, subscriberToId, time.Now()).
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
