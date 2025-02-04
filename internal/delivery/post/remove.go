package handlers

import (
	"context"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
)

func (h *PostGRPCHandler) RemovePost(ctx context.Context, req *postpb.RemovePostRequest) (*postpb.RemovePostResponse, error) {
	err := h.psP.Remove(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &postpb.RemovePostResponse{Success: true}, nil
}
