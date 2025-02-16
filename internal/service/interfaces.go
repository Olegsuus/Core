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
	ServiceGetSubscribers(ctx context.Context, userID string, settings models.GetManySettings) ([]models.User, error)
}

type PostServiceProvider interface {
	ServiceAdd(ctx context.Context, title, content, userID string) (string, error)
	ServiceRemove(ctx context.Context, id string) error
	ServiceGetMany(ctx context.Context, settings models.GetManySettings) ([]models.Post, error)
	ServiceGetFeed(ctx context.Context, subscriberID string, settings models.GetManySettings) ([]models.Post, error)
	ServiceGetPost(ctx context.Context, postID string) (models.Post, error)
}
