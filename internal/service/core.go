package service

import (
	storage "github.com/Olegsuus/Core/internal/storage/postgres"
)

type ServicesImpl struct {
	repository *storage.RepositoryImpl
}

func NewServicesImpl(repository *storage.RepositoryImpl) *ServicesImpl {
	return &ServicesImpl{
		repository: repository,
	}
}

func userEntityToModels(entity storage.UserEntity) *User {
	return &User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
	}
}

func postEntityToModels(entity storage.PostEntity) *Post {
	return &Post{
		ID:        entity.ID,
		Title:     entity.Title,
		Content:   entity.Content,
		UserID:    entity.UserID,
		CreatedAt: entity.CreatedAt,
	}
}
