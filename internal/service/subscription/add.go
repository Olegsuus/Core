package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
)

func (s *SubscriptionService) Subscribe(ctx context.Context, userID, subscribedToID string) error {
	subscription := &models.Subscription{
		SubscriberID:   userID,
		SubscribedToID: subscribedToID,
	}

	if err := s.subscriptionStorage.Subscribe(ctx, modelsToEntity(subscription)); err != nil {
		return fmt.Errorf("Storage.Subscribe: %w", err)
	}

	return nil
}
