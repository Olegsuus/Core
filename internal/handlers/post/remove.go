package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *PostGRPCHandler) RemovePost(ctx context.Context, req *postpb.RemovePostRequest) (*postpb.RemovePostResponse, error) {
	err := h.postService.RemovePost(ctx, req.GetId())
	if err != nil {
		h.l.Debug("ошибка при удалении нового поста", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, err
	}
	return &postpb.RemovePostResponse{Success: true}, nil
}
