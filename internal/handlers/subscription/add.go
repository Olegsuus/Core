package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *SubscriptionGRPCHandler) Subscribe(ctx context.Context, req *postpb.SubscribeRequest) (*postpb.SubscribeResponse, error) {
	if err := h.ssp.ServiceSubscribe(ctx, req.GetSubscriberId(), req.GetSubscribedToId()); err != nil {
		h.l.Debug("ошибка при подписке", slog.String("error:", fmt.Sprintf("%w", err)))
		return nil, err
	}

	return &postpb.SubscribeResponse{
		Success: true,
	}, nil
}
