package storage

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperrors "github.com/Olegsuus/Core/pkg/errors"
	"log/slog"
)

func (s *RepositoryImpl) Subscribe(ctx context.Context, param SubscribersParam) error {
	s.l.Info("Создание новой подписки", slog.String("details",
		fmt.Sprintf("пользователь: %s, подписывается на %s", param.SubscriberID, param.SubscribedToID)))

	subscriptionEntity := SubscriptionEntity{
		SubscriberID:   param.SubscriberID,
		SubscribedToID: param.SubscribedToID,
	}

	query, args := squirrel.
		Insert("subscriptions").
		Columns("subscriber_id", "subscribed_to_id").
		Values(subscriptionEntity.SubscriberID, subscriptionEntity.SubscribedToID).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	_, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при подписке",
		}
	}

	return nil
}

func (s *RepositoryImpl) GetSubscribers(ctx context.Context, userID string, settings GetManyParam) ([]UserEntity, error) {

	query, args := squirrel.
		Select("u.id", "u.name", "u.email", "u.created_at").
		From("subscriptions s").
		Join("users u ON s.subscriber_id = u.id").
		Where(squirrel.Eq{"s.subscribed_to_id": userID}).
		OrderBy("s.created_at DESC").
		Limit(uint64(settings.Limit)).
		Offset(uint64(settings.Offset)).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	var usersEntity []UserEntity
	if err := s.db.SelectContext(ctx, &usersEntity, query, args...); err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить подписчиков",
			Status:        404,
		}
	}

	return usersEntity, nil
}

func (s *RepositoryImpl) Unsubscribe(ctx context.Context, param SubscribersParam) error {
	s.l.Info("Создание новой подписки", slog.String("details",
		fmt.Sprintf("пользователь: %s, отписывается от %s", param.SubscriberID,
			param.SubscribedToID)))

	subscriptionEntity := SubscriptionEntity{
		SubscriberID:   param.SubscriberID,
		SubscribedToID: param.SubscribedToID,
	}

	query, args := squirrel.
		Delete("subscriptions").
		Where(squirrel.Eq{"subscriber_id": subscriptionEntity.SubscriberID, "subscribed_to_id": subscriptionEntity.SubscribedToID}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при отписке",
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось определить результат отписки",
		}
	}
	if rowsAffected == 0 {
		return apperrors.AppError{
			BusinessError: fmt.Sprintf("не найдена подписка для subscriber_id=%s и subscribed_to_id=%s",
				subscriptionEntity.SubscriberID, subscriptionEntity.SubscribedToID),
			UserError: "подписка не найдена",
			Status:    404,
		}
	}

	return nil
}
