package storage

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	"github.com/Olegsuus/Core/pkg/errors"
	"github.com/Olegsuus/Core/pkg/utils"
)

func (s *PostStorage) GetPost(ctx context.Context, postID string) (*models.Post, error) {
	query, args, err := squirrel.
		Select("id", "user_id", "title", "content", "created_at").
		From("posts").
		Where(squirrel.Eq{"id": postID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса на получение поста",
		}
	}

	var postEntity storage.PostEntity
	if err = s.db.GetContext(ctx, &postEntity, query, args...); err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить пост",
			Status:        404,
		}
	}

	return postEntityToModels(postEntity), nil

}

func (s *PostStorage) GetFeed(ctx context.Context, subscriberID string, limit, offset int) ([]*models.Post, error) {
	query, args, err := squirrel.
		Select("p.id", "p.user_id", "p.title", "p.content", "p.created_at").
		From("posts p").
		Join("subscriptions s ON p.user_id = s.subscribed_to_id").
		Where(squirrel.Eq{"s.subscriber_id": subscriberID}).
		OrderBy("p.created_at DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при составлении запроса на получении ленты",
			Status:        500,
		}
	}

	var postsEntity []storage.PostEntity
	if err = s.db.SelectContext(ctx, &postsEntity, query, args...); err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить ленту",
			Status:        400,
		}
	}

	posts := utils.MapAsync(postsEntity, postEntityToModels)

	return posts, nil
}

func (s *PostStorage) GetManyPosts(ctx context.Context, limit, offset int, sort string) ([]*models.Post, error) {
	query, args, err := squirrel.
		Select("id", "user_id", "title", "content", "created_at").
		From("posts").
		OrderBy("created_at " + sort).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при получении списка постов",
		}
	}

	var postsEntity []storage.PostEntity
	if err := s.db.SelectContext(ctx, &postsEntity, query, args...); err != nil {
		return nil, errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить список постов",
			Status:        404,
		}
	}

	posts := utils.MapAsync(postsEntity, postEntityToModels)

	return posts, nil
}
