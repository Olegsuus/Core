package service

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
)

func (s *UserService) GetUser(ctx context.Context, userID string) (*postpb.User, error) {
	user, err := s.userStorage.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("StorageGetUser: %w", err)
	}

	return modelsToGRPC(user), nil
}
