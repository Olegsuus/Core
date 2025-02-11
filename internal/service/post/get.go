package service

import (
	"context"
	"fmt"
	models2 "github.com/Olegsuus/Core/internal/models"
)

func (s *PostService) ServiceGetMany(ctx context.Context, settings models2.GetManyPostSettings) ([]models2.Post, error) {
	if settings.Limit <= 0 {
		settings.Limit = 15
	}

	if settings.Offset < 0 {
		settings.Offset = 0
	}

	posts, err := s.psP.StorageGetMany(ctx, settings)
	if err != nil {
		return nil, fmt.Errorf("StorageGetMany: %w", err)
	}

	return posts, nil
}

func (s *PostService) ServiceGetFeed(ctx context.Context, subscriberID string, settings models2.GetManyPostSettings) ([]models2.Post, error) {
	posts, err := s.psP.StorageGetFeed(ctx, subscriberID, settings)
	if err != nil {
		return nil, fmt.Errorf("StorageGetFeed: %w", err)
	}

	return posts, nil
}
