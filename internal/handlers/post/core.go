package handlers

import (
	"github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

type PostGRPCHandler struct {
	postpb.UnimplementedPostServiceServer
	postService service.PostService
	l           *slog.Logger
}

func NewPostGRPCHandler(postService service.PostService, l *slog.Logger) *PostGRPCHandler {
	return &PostGRPCHandler{
		postService: postService,
		l:           l,
	}
}
