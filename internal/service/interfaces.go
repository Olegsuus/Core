package service

import (
	"context"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
)

type UserService interface {
	AddUser(ctx context.Context, name, email, password string) (*postpb.User, error)
	GetUser(ctx context.Context, userID string) (*postpb.User, error)
}

type SubscriptionService interface {
	Subscribe(ctx context.Context, userID, subscribedToID string) error
	Unsubscribe(ctx context.Context, userID, subscribedToID string) error
	GetSubscribers(ctx context.Context, userID string, limit, offset int) ([]*postpb.User, error)
}

type PostService interface {
	AddPost(ctx context.Context, title, content, userID string) (*postpb.Post, error)
	RemovePost(ctx context.Context, id string) error
	GetManyPosts(ctx context.Context, limit, offset int, order bool) ([]*postpb.Post, error)
	GetFeed(ctx context.Context, subscriberID string, limit, offset int) ([]*postpb.Post, error)
	GetPost(ctx context.Context, postID string) (*postpb.Post, error)
}
