package storage

import (
	"context"
)

type Repository interface {
	User() UserRepository
	Subscription() SubscriptionRepository
	Post() PostRepository
}

type UserRepository interface {
	AddUser(ctx context.Context, user AddUserParam) (*UserEntity, error)
	GetUser(ctx context.Context, userID string) (*UserEntity, error)
}

type SubscriptionRepository interface {
	Unsubscribe(ctx context.Context, param SubscribersParam) error
	Subscribe(ctx context.Context, param SubscribersParam) error
	GetSubscribers(ctx context.Context, userID string, settings GetManyParam) ([]UserEntity, error)
}

type PostRepository interface {
	AddPost(ctx context.Context, post AddPostParam) (*PostEntity, error)
	GetManyPosts(ctx context.Context, settings GetManyParam) ([]PostEntity, error)
	RemovePost(ctx context.Context, id string) error
	GetFeed(ctx context.Context, subscriberID string, settings GetManyParam) ([]PostEntity, error)
	GetPost(ctx context.Context, postID string) (*PostEntity, error)
}
