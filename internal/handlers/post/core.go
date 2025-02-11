package handlers

import (
	"github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

type PostGRPCHandler struct {
	postpb.UnimplementedPostServiceServer
	psP service.PostServiceProvider
	l   *slog.Logger
}

func RegisterNewPostGRPCHandler(psP service.PostServiceProvider, l *slog.Logger) *PostGRPCHandler {
	return &PostGRPCHandler{
		psP: psP,
		l:   l,
	}
}
