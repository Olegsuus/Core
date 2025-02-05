package handlers

import (
	"context"
	"github.com/Olegsuus/Core/internal/domain/dto"
	domain "github.com/Olegsuus/Core/internal/domain/post"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
)

type PostGRPCHandler struct {
	postpb.UnimplementedPostServiceServer
	psP PostServiceProvider
}

type PostServiceProvider interface {
	Add(ctx context.Context, title, content string) (string, error)
	GetMany(ctx context.Context, settings dto.GetManyPostSettings) ([]*domain.Post, error)
	Remove(ctx context.Context, id string) error
}

func RegisterPostGRPCHandler(psP PostServiceProvider) *PostGRPCHandler {
	return &PostGRPCHandler{
		psP: psP,
	}
}
