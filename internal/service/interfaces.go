package service

import (
	"context"
)

type Service interface {
	User() UserService
	Subscription() SubscriptionService
	Post() PostService
}

type UserService interface {
	AddUser(ctx context.Context, param AddUserParam) (*User, error)
	GetUser(ctx context.Context, userID string) (*User, error)
}

type SubscriptionService interface {
	Subscribe(ctx context.Context, param SubscribersParam) error
	Unsubscribe(ctx context.Context, param SubscribersParam) error
	GetSubscribers(ctx context.Context, userID string, param GetManyParam) ([]User, error)
}

type PostService interface {
	AddPost(ctx context.Context, param AddPostParam) (*Post, error)
	RemovePost(ctx context.Context, id string) error
	GetManyPosts(ctx context.Context, settings GetManyParam) ([]Post, error)
	GetFeed(ctx context.Context, subscriberID string, settings GetManyParam) ([]Post, error)
	GetPost(ctx context.Context, postID string) (*Post, error)
}
