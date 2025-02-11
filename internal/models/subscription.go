package models

import "time"

type Subscription struct {
	ID             string    `db:"id" json:"id"`
	SubscriberID   string    `db:"subscriber_id" json:"subscriber_id"`
	SubscribedToID string    `db:"subscribed_to_id" json:"subscribed_to_id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
