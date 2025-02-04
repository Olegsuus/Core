package handlers

import (
	"context"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
)

func (h *PostGRPCHandler) AddPost(ctx context.Context, req *postpb.AddPostRequest) (*postpb.AddPostResponse, error) {
	id, err := h.psP.Add(ctx, req.GetTitle(), req.GetContent())
	if err != nil {
		return nil, err
	}
	return &postpb.AddPostResponse{Id: id}, nil
}
