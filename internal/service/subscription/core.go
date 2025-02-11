package service

import "github.com/Olegsuus/Core/internal/storage"

type SubscriptionService struct {
	ssp storage.SubscriptionStorageProvider
}

func RegisterNewSubscriptionService(ssp storage.SubscriptionStorageProvider) *SubscriptionService {
	return &SubscriptionService{
		ssp: ssp,
	}
}
