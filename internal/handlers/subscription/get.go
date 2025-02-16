package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *SubscriptionGRPCHandler) GetSubscribers(ctx context.Context, req *postpb.GetSubscribersRequest) (*postpb.GetSubscribersResponse, error) {
	userID := req.GetUserId()
	limit := req.GetLimit()
	page := req.GetPage()

	subscribers, err := h.subscriptionService.GetSubscribers(ctx, userID, int(limit), int(page))
	if err != nil {
		h.l.Debug("ошибка при получении подписчиков пользователя",
			slog.String("details", fmt.Sprintf("userID: %s", userID)),
			slog.String("error", fmt.Sprintf("%s", err)))

		return nil, err
	}

	return &postpb.GetSubscribersResponse{
		User: subscribers,
	}, nil
}
