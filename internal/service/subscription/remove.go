package service

import (
	"context"
	"fmt"
)

func (s *SubscriptionService) ServiceUnsubscribe(ctx context.Context, userID, subscribedToID string) error {
	if err := s.ssp.StorageUnsubscribe(ctx, userID, subscribedToID); err != nil {
		return fmt.Errorf("StorageUnsubscribe: %w", err)
	}
	return nil
}
