package service

import (
	"context"
	"fmt"
)

func (s *PostService) RemovePost(ctx context.Context, id string) error {

	if err := s.postStorage.RemovePost(ctx, id); err != nil {
		return fmt.Errorf("Storage.Remove: %w", err)
	}

	return nil
}
