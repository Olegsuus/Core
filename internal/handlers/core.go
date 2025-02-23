package handlers

import (
	"github.com/Olegsuus/Core/internal/service"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
)

type GRPCHandlers struct {
	service *service.ServicesImpl
	l       *slog.Logger
	postpb.UnimplementedPostServiceServer
	postpb.UnimplementedSubscriptionServiceServer
	postpb.UnimplementedUserServiceServer
}

func NewGRPCHandlers(service *service.ServicesImpl, l *slog.Logger) *GRPCHandlers {
	return &GRPCHandlers{
		service: service,
		l:       l,
	}
}

func userModelsToGRPC(user service.User) *postpb.User {
	return &postpb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}

func postModelsToGRPC(post service.Post) *postpb.Post {
	return &postpb.Post{
		Id:        post.ID,
		Title:     post.Title,
		UserId:    post.UserID,
		Content:   post.Content,
		CreatedAt: timestamppb.New(post.CreatedAt),
	}
}
