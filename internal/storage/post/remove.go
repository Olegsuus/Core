package storage

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/errors"
	"log"
)

func (s *PostStorage) Remove(ctx context.Context, id int64) error {
	const op = "storage.Remove"

	query := `DELETE FROM posts WHERE id = $1`

	ct, err := s.pg.Exec(ctx, query, id)
	if err != nil {
		log.Printf("ошибка при удалении поста: (%s: %w)", op, err)
		return errors.AppError{
			BusinessError: err.Error(),
			UserError:     "ошибка при удалении поста",
			Status:        500,
		}
	}
	if ct.RowsAffected() == 0 {
		log.Printf("не найден пост с таким id: %d", id)
		return errors.AppError{
			BusinessError: fmt.Sprintf("не найден пост с id = %d", id),
			UserError:     "пост не найден",
			Status:        400,
		}
	}

	log.Printf("пост успешно удален")

	return nil
}
