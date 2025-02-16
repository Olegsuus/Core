package storage

import (
	"context"
	"github.com/Olegsuus/Core/internal/models"
)

type UserStorageProvider interface {
	StorageAddUser(ctx context.Context, user *models.User) (string, error)
	StorageGetUser(ctx context.Context, userID string) (models.User, error)
}

type SubscriptionStorageProvider interface {
	StorageUnsubscribe(ctx context.Context, userID, subscribedToID string) error
	StorageAddSubscribe(ctx context.Context, userID, subscriberToId string) error
	StorageGetSubscribers(ctx context.Context, userID string, settings models.GetManySettings) ([]models.User, error)
}

type PostStorageProvider interface {
	StorageAddPost(ctx context.Context, post *models.Post) (string, error)
	StorageGetMany(ctx context.Context, settings models.GetManySettings) ([]models.Post, error)
	StorageRemovePost(ctx context.Context, id string) error
	StorageGetFeed(ctx context.Context, subscriberID string, settings models.GetManySettings) ([]models.Post, error)
	StorageGetPost(ctx context.Context, postID string) (models.Post, error)
}
