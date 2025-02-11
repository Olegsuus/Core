package handlers

import (
	"github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

type UserGRPCHandler struct {
	postpb.UnimplementedUserServiceServer
	usp service.UserServiceProvider
	l   *slog.Logger
}

func RegisterNewUserGRPCHandler(usp service.UserServiceProvider, l *slog.Logger) *UserGRPCHandler {
	return &UserGRPCHandler{
		usp: usp,
		l:   l,
	}
}
