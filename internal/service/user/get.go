package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
)

func (s *UserService) ServiceGet(ctx context.Context, userID string) (models.User, error) {
	user, err := s.suP.StorageGetUser(ctx, userID)
	if err != nil {
		return models.User{}, fmt.Errorf("StorageGetUser: %w", err)
	}

	return user, nil
}
