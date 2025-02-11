package service

import (
	"context"
	models2 "github.com/Olegsuus/Core/internal/models"
)

type UserServiceProvider interface {
	ServiceAdd(ctx context.Context, name, email, password string) (string, error)
}

type SubscriptionServiceProvider interface {
	ServiceSubscribe(ctx context.Context, userID, subscribedToID string) error
	ServiceUnsubscribe(ctx context.Context, userID, subscribedToID string) error
}

type PostServiceProvider interface {
	ServiceAdd(ctx context.Context, title, content string) (string, error)
	ServiceRemove(ctx context.Context, id string) error
	ServiceGetMany(ctx context.Context, settings models2.GetManyPostSettings) ([]models2.Post, error)
	ServiceGetFeed(ctx context.Context, subscriberID string, settings models2.GetManyPostSettings) ([]models2.Post, error)
}
