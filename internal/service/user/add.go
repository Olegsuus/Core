package service

import (
	"context"
	"fmt"
	"github.com/Olegsuus/Core/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) ServiceAdd(ctx context.Context, name, email, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	id, err := s.suP.StorageAddUser(ctx, user)
	if err != nil {
		return "", fmt.Errorf("StorageAdd: %w", err)
	}

	return id, nil
}
