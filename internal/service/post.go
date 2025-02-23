package service

import (
	"context"
	"fmt"
	storage "github.com/Olegsuus/Core/internal/storage/postgres"
	"time"
)

func (s *ServicesImpl) AddPost(ctx context.Context, dto AddPostParam) (*Post, error) {
	post, err := s.repository.AddPost(ctx, storage.AddPostParam{
		ID:        dto.ID,
		UserID:    dto.UserID,
		Title:     dto.Title,
		Content:   dto.Content,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return nil, fmt.Errorf("StorageAdd: %w", err)
	}

	return postEntityToModels(*post), nil
}

func (s *ServicesImpl) GetPost(ctx context.Context, postID string) (*Post, error) {
	post, err := s.repository.GetPost(ctx, postID)
	if err != nil {
		return nil, fmt.Errorf("Storage.GetPost: %w", err)
	}

	return postEntityToModels(*post), nil
}

func (s *ServicesImpl) GetFeed(ctx context.Context, subscriberID string, settings GetManyParam) ([]Post, error) {
	if settings.Limit <= 0 {
		settings.Limit = 15
	}

	if settings.Offset < 0 {
		settings.Offset = 0
	}

	posts, err := s.repository.GetFeed(ctx, subscriberID, storage.GetManyParam{
		Limit:  settings.Limit,
		Offset: settings.Offset,
	})

	if err != nil {
		return nil, fmt.Errorf("Storage.GetFeed: %w", err)
	}

	var modelsPosts []Post
	for _, post := range posts {
		modelsPosts = append(modelsPosts, *postEntityToModels(post))
	}

	return modelsPosts, nil
}

func (s *ServicesImpl) GetManyPosts(ctx context.Context, settings GetManyParam) ([]Post, error) {
	if settings.Limit <= 0 {
		settings.Limit = 15
	}

	if settings.Offset < 0 {
		settings.Offset = 0
	}

	var sort string
	if settings.Order {
		sort = "DESC"
	} else {
		sort = "ASC"
	}

	posts, err := s.repository.GetManyPosts(ctx, storage.GetManyParam{
		Limit:  settings.Limit,
		Offset: settings.Offset,
		Sort:   sort,
	})

	if err != nil {
		return nil, fmt.Errorf("Storage.GetManyPosts: %w", err)
	}

	var modelsPosts []Post
	for _, post := range posts {
		modelsPosts = append(modelsPosts, *postEntityToModels(post))
	}

	return modelsPosts, nil
}

func (s *ServicesImpl) RemovePost(ctx context.Context, id string) error {
	if err := s.repository.RemovePost(ctx, id); err != nil {
		return fmt.Errorf("Storage.Remove: %w", err)
	}

	return nil
}
