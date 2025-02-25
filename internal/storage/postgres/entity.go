package storage

import "time"

type UserEntity struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
}

type PostEntity struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type SubscriptionEntity struct {
	ID             string    `db:"id"`
	SubscriberID   string    `db:"subscriber_id"`
	SubscribedToID string    `db:"subscribed_to_id"`
	CreatedAt      time.Time `db:"created_at"`
}

type GetManyParam struct {
	Limit  int
	Offset int
	Sort   string
}

type AddUserParam struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type AddPostParam struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	CreatedAt time.Time
}

type SubscribersParam struct {
	SubscriberID   string
	SubscribedToID string
}
