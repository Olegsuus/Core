package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
)

func (s *SubscriptionService) ServiceGetSubscribers(ctx context.Context, userID string, settings models.GetManySettings,
) ([]models.User, error) {

	var subscribers []models.User

	_, err := s.usp.StorageGetUser(ctx, userID)
	if err != nil {
		return subscribers, fmt.Errorf("StorageGetUser: %w", err)
	}

	subscribers, err = s.ssp.StorageGetSubscribers(ctx, userID, settings)
	if err != nil {
		return subscribers, fmt.Errorf("StorageGetUser: %w", err)
	}

	return subscribers, nil
}
