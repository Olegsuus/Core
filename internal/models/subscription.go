package models

import "time"

type Subscription struct {
	ID             string
	SubscriberID   string
	SubscribedToID string
	CreatedAt      time.Time
}
