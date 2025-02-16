package handlers

import (
	"github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

type UserGRPCHandler struct {
	postpb.UnimplementedUserServiceServer
	userService service.UserService
	l           *slog.Logger
}

func NewUserGRPCHandler(userService service.UserService, l *slog.Logger) *UserGRPCHandler {
	return &UserGRPCHandler{
		userService: userService,
		l:           l,
	}
}
