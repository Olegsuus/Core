package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *SubscriptionGRPCHandler) Unsubscribe(ctx context.Context, req *postpb.UnsubscribeRequest) (*postpb.UnsubscribeResponse, error) {
	if err := h.subscriptionService.Unsubscribe(ctx, req.GetSubscriberId(), req.GetSubscribedToId()); err != nil {
		h.l.Debug("ошибка при отписке", slog.String("error:", fmt.Sprintf("%s", err)))
		return nil, err
	}

	return &postpb.UnsubscribeResponse{
		Success: true,
	}, nil
}
