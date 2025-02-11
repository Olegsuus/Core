package service

import (
	"github.com/Olegsuus/Core/internal/storage"
)

type UserService struct {
	suP storage.UserStorageProvider
}

func RegisterNewServiceUser(suP storage.UserStorageProvider) *UserService {
	return &UserService{
		suP: suP,
	}
}
