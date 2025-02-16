package models

import "time"

type Post struct {
	ID        string
	UserID    string
	Title     string
	Content   string
	CreatedAt time.Time
}
