package service

import "github.com/Olegsuus/Core/internal/storage"

type SubscriptionService struct {
	ssp storage.SubscriptionStorageProvider
	usp storage.UserStorageProvider
}

func RegisterNewSubscriptionService(ssp storage.SubscriptionStorageProvider, usp storage.UserStorageProvider) *SubscriptionService {
	return &SubscriptionService{
		ssp: ssp,
		usp: usp,
	}
}
