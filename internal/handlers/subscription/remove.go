package handlers

import (
	"context"
	"fmt"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

func (h *SubscriptionGRPCHandler) Unsubscribe(ctx context.Context, req *postpb.UnsubscribeRequest) (*postpb.UnsubscribeResponse, error) {
	if err := h.ssp.ServiceUnsubscribe(ctx, req.GetSubscriberId(), req.GetSubscribedToId()); err != nil {
		h.l.Debug("ошибка при отписке", slog.String("error:", fmt.Sprintf("%w", err)))
		return nil, err
	}

	return &postpb.UnsubscribeResponse{
		Success: true,
	}, nil
}
