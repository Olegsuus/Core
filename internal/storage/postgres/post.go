package storage

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	apperrors "github.com/Olegsuus/Core/pkg/errors"
)

func (s *RepositoryImpl) AddPost(ctx context.Context, postParam AddPostParam) (*PostEntity, error) {
	s.l.Info("создание нового поста:", postParam)

	postEntity := PostEntity{
		Title:   postParam.Title,
		Content: postParam.Content,
		UserID:  postParam.UserID,
	}

	query, args := squirrel.
		Insert("posts").
		Columns("user_id", "title", "content").
		Values(postEntity.UserID, postEntity.Title, postEntity.Content).
		Suffix("RETURNING id, user_id, title, content, created_at").
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	if err := s.db.GetContext(ctx, &postEntity, query, args...); err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при сохранении нового поста",
		}
	}

	return &postEntity, nil
}

func (s *RepositoryImpl) GetPost(ctx context.Context, postID string) (*PostEntity, error) {
	query, args := squirrel.
		Select("id", "user_id", "title", "content", "created_at").
		From("posts").
		Where(squirrel.Eq{"id": postID}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	var postEntity PostEntity
	if err := s.db.GetContext(ctx, &postEntity, query, args...); err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить пост",
			Status:        404,
		}
	}

	return &postEntity, nil

}

func (s *RepositoryImpl) GetFeed(ctx context.Context, subscriberID string, settings GetManyParam) ([]PostEntity, error) {
	query, args := squirrel.
		Select("p.id", "p.user_id", "p.title", "p.content", "p.created_at").
		From("posts p").
		Join("subscriptions s ON p.user_id = s.subscribed_to_id").
		Where(squirrel.Eq{"s.subscriber_id": subscriberID}).
		OrderBy("p.created_at DESC").
		Limit(uint64(settings.Limit)).
		Offset(uint64(settings.Offset)).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	var postsEntity []PostEntity
	if err := s.db.SelectContext(ctx, &postsEntity, query, args...); err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить ленту",
			Status:        400,
		}
	}

	return postsEntity, nil
}

func (s *RepositoryImpl) GetManyPosts(ctx context.Context, settings GetManyParam) ([]PostEntity, error) {
	query, args := squirrel.
		Select("id", "user_id", "title", "content", "created_at").
		From("posts").
		OrderBy("created_at " + settings.Sort).
		Limit(uint64(settings.Limit)).
		Offset(uint64(settings.Offset)).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	var postsEntity []PostEntity
	if err := s.db.SelectContext(ctx, &postsEntity, query, args...); err != nil {
		return nil, apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось получить список постов",
			Status:        404,
		}
	}

	return postsEntity, nil
}

func (s *RepositoryImpl) RemovePost(ctx context.Context, id string) error {
	s.l.Info("удаление поста", "id", id)

	query, args := squirrel.
		Delete("posts").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при удалении поста",
			Status:        500,
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return apperrors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось определить количество удаленных записей",
			Status:        500,
		}
	}

	if rowsAffected == 0 {
		return apperrors.AppError{
			BusinessError: fmt.Sprintf("не найден пост с id = %s", id),
			UserError:     "пост не найден",
			Status:        404,
		}
	}

	return nil
}
