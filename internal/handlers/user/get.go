package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *UserGRPCHandler) GetUser(ctx context.Context, req *postpb.GetUserRequest) (*postpb.GetUserResponse, error) {
	user, err := h.userService.GetUser(ctx, req.GetId())
	if err != nil {
		h.l.Debug("ошибка при получении пользователя", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, err
	}

	return &postpb.GetUserResponse{
		User: user,
	}, nil
}
