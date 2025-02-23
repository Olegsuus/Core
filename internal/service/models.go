package service

import "time"

type Post struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	CreatedAt time.Time
}

type AddPostParam struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	CreatedAt time.Time
}

type Subscription struct {
	ID             string
	SubscriberID   string
	SubscribedToID string
	CreatedAt      time.Time
}

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type AddUserParam struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type GetManyParam struct {
	Limit  int
	Offset int
	Order  bool
}

type SubscribersParam struct {
	SubscriberID   string
	SubscribedToID string
}
