package service

import (
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostService struct {
	postStorage storage.PostStorage
}

func NewPostService(postStorage storage.PostStorage) *PostService {
	return &PostService{
		postStorage: postStorage,
	}
}

func modelsToEntity(post *models.Post) *storage.PostEntity {
	return &storage.PostEntity{
		ID:        post.ID,
		UserID:    post.UserID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
	}
}

func modelsToGRPC(post *models.Post) *postpb.Post {
	return &postpb.Post{
		Id:        post.ID,
		UserId:    post.UserID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: timestamppb.New(post.CreatedAt),
	}
}
