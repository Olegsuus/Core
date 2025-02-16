package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *UserGRPCHandler) AddUser(ctx context.Context, req *postpb.AddUserRequest) (*postpb.AddUserResponse, error) {
	user, err := h.userService.AddUser(ctx, req.GetName(), req.GetEmail(), req.GetPassword())
	if err != nil {
		h.l.Debug("ошибка при добавлении нового пользователя", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, err
	}

	return &postpb.AddUserResponse{
		Id: user.Id,
	}, nil
}
