package handlers

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *GRPCHandlers) Subscribe(ctx context.Context, req *postpb.SubscribeRequest) (*postpb.SubscribeResponse, error) {
	param := service.SubscribersParam{
		SubscriberID:   req.GetSubscriberId(),
		SubscribedToID: req.GetSubscribedToId(),
	}

	if err := h.service.Subscribe(ctx, param); err != nil {
		h.l.Error("ошибка при подписке", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, err
	}

	return &postpb.SubscribeResponse{
		Success: true,
	}, nil
}

func (h *GRPCHandlers) GetSubscribers(ctx context.Context, req *postpb.GetSubscribersRequest) (*postpb.GetSubscribersResponse, error) {
	userID := req.GetUserId()
	limit := req.GetLimit()
	page := req.GetPage()
	offset := (page - 1) * limit

	subscribers, err := h.service.GetSubscribers(ctx, userID, service.GetManyParam{
		Limit:  int(limit),
		Offset: int(offset),
	})
	if err != nil {
		h.l.Error("ошибка при получении подписчиков пользователя",
			slog.String("details", fmt.Sprintf("userID: %s", userID)),
			slog.String("error", fmt.Sprintf("%s", err)))

		return nil, fmt.Errorf("Serivce.GetSubscribers: %w", err)
	}

	var subscribersPB []*postpb.User
	for _, sub := range subscribers {
		subscribersPB = append(subscribersPB, userModelsToGRPC(sub))
	}

	return &postpb.GetSubscribersResponse{
		User: subscribersPB,
	}, nil
}

func (h *GRPCHandlers) Unsubscribe(ctx context.Context, req *postpb.UnsubscribeRequest) (*postpb.UnsubscribeResponse, error) {
	param := service.SubscribersParam{
		SubscriberID:   req.GetSubscriberId(),
		SubscribedToID: req.GetSubscribedToId(),
	}

	if err := h.service.Unsubscribe(ctx, param); err != nil {
		h.l.Error("ошибка при отписке", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, fmt.Errorf("Service.Unsubscribe: %w", err)
	}

	return &postpb.UnsubscribeResponse{
		Success: true,
	}, nil
}
