package service

import (
	"context"
	"fmt"
	storage "github.com/Olegsuus/Core/internal/storage/postgres"
)

func (s *ServicesImpl) Subscribe(ctx context.Context, param SubscribersParam) error {
	subscription := storage.SubscribersParam{
		SubscriberID:   param.SubscriberID,
		SubscribedToID: param.SubscribedToID,
	}

	if err := s.repository.Subscribe(ctx, subscription); err != nil {
		return fmt.Errorf("Storage.Subscribe: %w", err)
	}

	return nil
}

func (s *ServicesImpl) GetSubscribers(ctx context.Context, userID string, settings GetManyParam) ([]User, error) {

	if settings.Limit <= 0 {
		settings.Limit = 15
	}

	if settings.Offset < 0 {
		settings.Offset = 0
	}

	subscribers, err := s.repository.GetSubscribers(ctx, userID, storage.GetManyParam{
		Limit:  settings.Limit,
		Offset: settings.Offset,
	})

	if err != nil {
		return nil, fmt.Errorf("Storage.GetSubscribers: %w", err)
	}

	var modelsSubscribers []User
	for _, sub := range subscribers {
		modelsSubscribers = append(modelsSubscribers, *userEntityToModels(sub))
	}

	return modelsSubscribers, nil
}

func (s *ServicesImpl) Unsubscribe(ctx context.Context, param SubscribersParam) error {
	subscription := storage.SubscribersParam{
		SubscribedToID: param.SubscribedToID,
		SubscriberID:   param.SubscriberID,
	}

	if err := s.repository.Unsubscribe(ctx, subscription); err != nil {
		return fmt.Errorf("Storage.Unsubscribe: %w", err)
	}

	return nil
}
