package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *PostGRPCHandler) AddPost(ctx context.Context, req *postpb.AddPostRequest) (*postpb.AddPostResponse, error) {
	id, err := h.psP.ServiceAdd(ctx, req.GetTitle(), req.GetContent(), req.GetUserId())
	if err != nil {
		h.l.Debug("ошибка при добавлении нового поста", slog.String("error:", fmt.Sprintf("%w", err)))
		return nil, err
	}
	return &postpb.AddPostResponse{Id: id}, nil
}
