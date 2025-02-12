package service

import (
	"github.com/Olegsuus/Core/internal/storage"
)

type PostService struct {
	psP storage.PostStorageProvider
	usP storage.UserStorageProvider
}

func RegisterPostService(psP storage.PostStorageProvider, usP storage.UserStorageProvider) *PostService {
	return &PostService{
		psP: psP,
		usP: usP,
	}
}
