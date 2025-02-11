package service

import (
	"github.com/Olegsuus/Core/internal/storage"
)

type PostService struct {
	psP storage.PostStorageProvider
}

func RegisterPostService(psP storage.PostStorageProvider) *PostService {
	return &PostService{
		psP: psP,
	}
}
