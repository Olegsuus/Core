package handlers

import (
	"context"
	"fmt"
	models "github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
	"time"
)

func (h *GRPCHandlers) AddUser(ctx context.Context, req *postpb.AddUserRequest) (*postpb.AddUserResponse, error) {
	user, err := h.service.AddUser(ctx, models.AddUserParam{
		ID:        "",
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		Password:  req.GetPassword(),
		CreatedAt: time.Now(),
	})

	if err != nil {
		h.l.Error("ошибка при добавлении нового пользователя", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, fmt.Errorf("Service.AddUser: %w", err)
	}

	return &postpb.AddUserResponse{
		Id: user.ID,
	}, nil
}

func (h *GRPCHandlers) GetUser(ctx context.Context, req *postpb.GetUserRequest) (*postpb.GetUserResponse, error) {
	user, err := h.service.GetUser(ctx, req.GetId())
	if err != nil {
		h.l.Error("ошибка при получении пользователя", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, fmt.Errorf("Service.GetUser: %w", err)
	}

	return &postpb.GetUserResponse{
		User: userModelsToGRPC(*user),
	}, nil
}
