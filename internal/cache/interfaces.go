package cache

import "github.com/Olegsuus/Core/internal/service"

type PostCache interface {
	GetFromCache(id string) (*service.Post, bool)
	AddToCache(id string, post *service.Post)
	RemoveFromCache(id string)
}
