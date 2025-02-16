package handlers

import (
	"github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

type SubscriptionGRPCHandler struct {
	postpb.UnimplementedSubscriptionServiceServer
	subscriptionService service.SubscriptionService
	l                   *slog.Logger
}

func NewSubscriptionGRPCHandler(subscriptionService service.SubscriptionService, l *slog.Logger) *SubscriptionGRPCHandler {
	return &SubscriptionGRPCHandler{
		subscriptionService: subscriptionService,
		l:                   l,
	}
}
