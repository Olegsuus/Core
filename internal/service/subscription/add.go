package service

import (
	"context"
	"fmt"
)

func (s *SubscriptionService) ServiceSubscribe(ctx context.Context, userID, subscribedToID string) error {
	_, err := s.usp.StorageGetUser(ctx, subscribedToID)
	if err != nil {
		return fmt.Errorf("StorageGetUser: %w", err)
	}

	if err := s.ssp.StorageAddSubscribe(ctx, userID, subscribedToID); err != nil {
		return fmt.Errorf("StorageSubscribe: %w", err)
	}
	return nil
}
