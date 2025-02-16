package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
)

func (s *SubscriptionService) Unsubscribe(ctx context.Context, userID, subscribedToID string) error {
	subscription := &models.Subscription{
		SubscribedToID: subscribedToID,
		SubscriberID:   userID,
	}

	if err := s.subscriptionStorage.Unsubscribe(ctx, modelsToEntity(subscription)); err != nil {
		return fmt.Errorf("Storage.Unsubscribe: %w", err)
	}
	return nil
}
