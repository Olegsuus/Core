package handlers

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

func (h *SubscriptionGRPCHandler) GetSubscribers(ctx context.Context, req *postpb.GetSubscribersRequest) (*postpb.GetSubscribersResponse, error) {
	userID := req.GetUserId()
	limit := req.GetLimit()
	page := req.GetPage()
	offset := (page - 1) * limit

	setings := models.GetManySettings{
		Limit:  int(limit),
		Offset: int(offset),
	}

	subscribers, err := h.ssp.ServiceGetSubscribers(ctx, userID, setings)
	if err != nil {
		h.l.Debug("ошибка при получении подписчиков пользователя",
			slog.String("details", fmt.Sprintf("userID: %s", userID)),
			slog.String("error", fmt.Sprintf("%w", err)))

		return nil, err
	}

	var pbUsers []*postpb.User
	for _, sub := range subscribers {
		pbUsers = append(pbUsers, &postpb.User{
			Id:        sub.ID,
			Name:      sub.Name,
			Email:     sub.Email,
			CreatedAt: timestamppb.New(sub.CreatedAt),
		})
	}

	return &postpb.GetSubscribersResponse{
		User: pbUsers,
	}, nil
}
