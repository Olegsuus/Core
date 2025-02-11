package service

import (
	"context"
	"fmt"
)

func (s *PostService) ServiceRemove(ctx context.Context, id string) error {

	if err := s.psP.StorageRemovePost(ctx, id); err != nil {
		return fmt.Errorf("StorageRemove: %w", err)
	}

	return nil
}
