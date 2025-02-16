package service

import (
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserService struct {
	userStorage storage.UserStorage
}

func NewServiceUser(userStorage storage.UserStorage) *UserService {
	return &UserService{
		userStorage: userStorage,
	}
}

func modelsToEntity(user *models.User) *storage.UserEntity {
	return &storage.UserEntity{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func modelsToGRPC(user *models.User) *postpb.User {
	return &postpb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
