package service

import (
	"context"
	models "github.com/Olegsuus/Core/internal/models"
)

type UserServiceProvider interface {
	ServiceAdd(ctx context.Context, name, email, password string) (string, error)
	ServiceGet(ctx context.Context, userID string) (models.User, error)
}

type SubscriptionServiceProvider interface {
	ServiceSubscribe(ctx context.Context, userID, subscribedToID string) error
	ServiceUnsubscribe(ctx context.Context, userID, subscribedToID string) error
}

type PostServiceProvider interface {
	ServiceAdd(ctx context.Context, title, content, userID string) (string, error)
	ServiceRemove(ctx context.Context, id string) error
	ServiceGetMany(ctx context.Context, settings models.GetManyPostSettings) ([]models.Post, error)
	ServiceGetFeed(ctx context.Context, subscriberID string, settings models.GetManyPostSettings) ([]models.Post, error)
}
