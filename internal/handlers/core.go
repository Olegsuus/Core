package handlers

import (
	"github.com/Olegsuus/Core/internal/service"
	postpb2 "github.com/Olegsuus/Core/proto/gen/go/core/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

type GRPCHandlers struct {
	service *service.ServicesImpl
	l       *slog.Logger
	postpb2.UnimplementedPostServiceServer
	postpb2.UnimplementedSubscriptionServiceServer
	postpb2.UnimplementedUserServiceServer
}

func NewGRPCHandlers(service *service.ServicesImpl, l *slog.Logger) *GRPCHandlers {
	return &GRPCHandlers{
		service: service,
		l:       l,
	}
}

func userModelsToGRPC(user service.User) *postpb2.User {
	return &postpb2.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}

func postModelsToGRPC(post service.Post) *postpb2.Post {
	return &postpb2.Post{
		Id:        post.ID,
		Title:     post.Title,
		UserId:    post.UserID,
		Content:   post.Content,
		CreatedAt: timestamppb.New(post.CreatedAt),
	}
}
