package storage

import (
	"context"
	"github.com/Olegsuus/Core/internal/models"
)

type UserStorage interface {
	AddUser(ctx context.Context, entity UserEntity) (*models.User, error)
	GetUser(ctx context.Context, userID string) (*models.User, error)
}

type SubscriptionStorage interface {
	Unsubscribe(ctx context.Context, subscriptionEntity *SubscriptionEntity) error
	Subscribe(ctx context.Context, subscriptionEntity *SubscriptionEntity) error
	GetSubscribers(ctx context.Context, userID string, limit, offset int) ([]*models.User, error)
}

type PostStorage interface {
	AddPost(ctx context.Context, postEntity PostEntity) (*models.Post, error)
	GetManyPosts(ctx context.Context, limit, offset int, sort string) ([]*models.Post, error)
	RemovePost(ctx context.Context, id string) error
	GetFeed(ctx context.Context, subscriberID string, limit, offset int) ([]*models.Post, error)
	GetPost(ctx context.Context, postID string) (*models.Post, error)
}
