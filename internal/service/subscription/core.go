package service

import (
	"github.com/Olegsuus/Core/internal/models"
	"github.com/Olegsuus/Core/internal/storage"
	postpb "github.com/Olegsuus/Core/settings_grpc/go/core/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SubscriptionService struct {
	subscriptionStorage storage.SubscriptionStorage
}

func NewSubscriptionService(subscriptionStorage storage.SubscriptionStorage) *SubscriptionService {
	return &SubscriptionService{
		subscriptionStorage: subscriptionStorage,
	}
}

func modelsToEntity(subscription *models.Subscription) *storage.SubscriptionEntity {
	return &storage.SubscriptionEntity{
		SubscriberID:   subscription.SubscriberID,
		SubscribedToID: subscription.SubscribedToID,
	}
}

func modelsToGRPC(user *models.User) *postpb.User {
	return &postpb.User{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}
