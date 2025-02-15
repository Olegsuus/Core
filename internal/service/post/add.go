package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
)

func (s *PostService) ServiceAdd(ctx context.Context, title, content, userID string) (string, error) {
	user, err := s.usP.StorageGetUser(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("StorageGetUser: %w", err)
	}

	newPost := &models.Post{
		Title:   title,
		Content: content,
		UserID:  user.ID,
	}

	id, err := s.psP.StorageAddPost(ctx, newPost)
	if err != nil {
		return "", fmt.Errorf("StorageAdd: %w", err)
	}

	return id, nil
}
