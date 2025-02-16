package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
)

func (s *PostService) AddPost(ctx context.Context, title, content, userID string) (*postpb.Post, error) {
	newPost := &models.Post{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	post, err := s.postStorage.AddPost(ctx, modelsToEntity(newPost))
	if err != nil {
		return nil, fmt.Errorf("StorageAdd: %w", err)
	}

	return modelsToGRPC(post), nil
}
