package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/pkg/utils"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
)

func (s *PostService) GetPost(ctx context.Context, postID string) (*postpb.Post, error) {
	post, err := s.postStorage.GetPost(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("Storage.GetPost: %w", err)
	}

	return modelsToGRPC(post), nil
}

func (s *PostService) GetFeed(ctx context.Context, subscriberID string, limit, offset int) ([]*postpb.Post, error) {
	if limit <= 0 {
		limit = 15
	}

	if offset < 0 {
		offset = 0
	}

	posts, err := s.postStorage.GetFeed(ctx, subscriberID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Storage.GetFeed: %w", err)
	}

	pbPosts := utils.MapAsync(posts, modelsToGRPC)

	return pbPosts, nil
}

func (s *PostService) GetManyPosts(ctx context.Context, limit, offset int, order bool) ([]*postpb.Post, error) {
	if limit <= 0 {
		limit = 15
	}

	if offset < 0 {
		offset = 0
	}

	var sort string
	if order {
		sort = "DESC"
	} else {
		sort = "ASC"
	}

	posts, err := s.postStorage.GetManyPosts(ctx, limit, offset, sort)
	if err != nil {
		return nil, fmt.Errorf("Storage.GetManyPosts: %w", err)
	}

	pbPosts := utils.MapAsync(posts, modelsToGRPC)

	return pbPosts, nil
}
