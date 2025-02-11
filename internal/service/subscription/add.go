package service

import (
	"context"
	"fmt"
)

func (s *SubscriptionService) ServiceSubscribe(ctx context.Context, userID, subscribedToID string) error {
	if err := s.ssp.StorageAddSubscribe(ctx, userID, subscribedToID); err != nil {
		return fmt.Errorf("StorageSubscribe: %w", err)
	}
	return nil
}
