package service

import (
	"context"
	"fmt"
	storage "github.com/Olegsuus/Core/internal/storage/postgres"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *ServicesImpl) AddUser(ctx context.Context, userParam AddUserParam) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userParam.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user, err := s.repository.AddUser(ctx, storage.AddUserParam{
		ID:        userParam.ID,
		Name:      userParam.Name,
		Email:     userParam.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, fmt.Errorf("Storage.AddUser: %w", err)
	}

	return userEntityToModels(*user), nil
}

func (s *ServicesImpl) GetUser(ctx context.Context, userID string) (*User, error) {
	user, err := s.repository.GetUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("StorageGetUser: %w", err)
	}

	return userEntityToModels(*user), nil
}
