package storage

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/Olegsuus/Core/pkg/errors"
)

func (s *PostStorage) RemovePost(ctx context.Context, id string) error {
	s.l.Info("удаление поста", "id", id)

	query, args, err := squirrel.
		Delete("posts").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при удалении поста",
		}
	}

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при удалении поста",
			Status:        500,
		}
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.AppError{
			BusinessError: err.Error(),
			UserError:     "не удалось определить количество удаленных записей",
			Status:        500,
		}
	}

	if rowsAffected == 0 {
		return errors.AppError{
			BusinessError: fmt.Sprintf("не найден пост с id = %s", id),
			UserError:     "пост не найден",
			Status:        404,
		}
	}

	return nil
}
