package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (h *UserGRPCHandler) GetUser(ctx context.Context, req *postpb.GetUserRequest) (*postpb.GetUserResponse, error) {
	user, err := h.usp.ServiceGetUser(ctx, req.GetId())
	if err != nil {
		h.l.Debug("ошибка при получении пользователя", slog.String("error:", fmt.Sprintf("%w", err)))
	}

	pUser := &postpb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}

	return &postpb.GetUserResponse{
		User: pUser,
	}, nil
}
