package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/pkg/utils"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
)

func (s *SubscriptionService) GetSubscribers(ctx context.Context, userID string, limit, offset int,
) ([]*postpb.User, error) {

	subscribers, err := s.subscriptionStorage.GetSubscribers(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Storage.GetSubscribers: %w", err)
	}

	pbSubscribers := utils.MapAsync(subscribers, modelsToGRPC)

	return pbSubscribers, nil
}
