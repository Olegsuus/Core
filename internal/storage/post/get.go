package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	models2 "github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/pkg/errors"
)

func (s *PostStorage) StorageGetMany(ctx context.Context, settings models2.GetManyPostSettings) ([]models2.Post, error) {
	order := "ASC"
	if settings.SortDesc {
		order = "DESC"
	}

	query, args, err := squirrel.
		Select("id", "title", "content", "created_at").
		From("post").
		OrderBy("created_at" + order).
		Limit(uint64(settings.Limit)).
		Offset(uint64(settings.Offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при получении списка постов",
		}
	}

	var posts []models2.Post
	if err := s.db.SelectContext(ctx, &posts, query, args...); err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить список постов",
			Status:        400,
		}
	}

	return posts, nil
}

func (s *PostStorage) StorageGetFeed(ctx context.Context, subscriberID string, settings models2.GetManyPostSettings) ([]models2.Post, error) {

	query, args, err := squirrel.
		Select("p.id", "p.user_id", "p.title", "p.content", "p.created_at").
		From("posts p").
		Join("subscriptions s ON p.user_id = s.subscribed_to_id").
		Where(squirrel.Eq{"s.subscriber_id": subscriberID}).
		OrderBy("p.created_at " + "DESC").
		Limit(uint64(settings.Limit)).
		Offset(uint64(settings.Offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса на получении ленты",
			Status:        500,
		}
	}

	var posts []models2.Post
	if err = s.db.GetContext(ctx, &posts, query, args...); err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить ленту",
			Status:        400,
		}
	}

	return posts, nil
}
