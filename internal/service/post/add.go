package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
)

func (s *PostService) ServiceAdd(ctx context.Context, title, content string) (string, error) {
	newPost := &models.Post{
		Title:   title,
		Content: content,
	}

	id, err := s.psP.StorageAddPost(ctx, newPost)
	if err != nil {
		return "", fmt.Errorf("StorageAdd: %w", err)
	}

	return id, nil
}
