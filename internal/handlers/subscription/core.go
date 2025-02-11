package handlers

import (
	"github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"log/slog"
)

type SubscriptionGRPCHandler struct {
	postpb.UnimplementedSubscriptionServiceServer
	ssp service.SubscriptionServiceProvider
	l   *slog.Logger
}

func RegisterNewSubscriptionGRPCHandler(ssp service.SubscriptionServiceProvider, l *slog.Logger) *SubscriptionGRPCHandler {
	return &SubscriptionGRPCHandler{
		ssp: ssp,
		l:   l,
	}
}
