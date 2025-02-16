package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) AddUser(ctx context.Context, name, email, password string) (*postpb.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	user, err = s.userStorage.AddUser(ctx, modelsToEntity(user))
	if err != nil {
		return nil, fmt.Errorf("Storage.AddUser: %w", err)
	}

	return modelsToGRPC(user), nil
}
