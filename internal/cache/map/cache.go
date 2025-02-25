package cache

import (
	"github.com/Olegsuus/Core/internal/service"
	"sync"
)

type MapCache struct {
	mu    sync.RWMutex
	cache map[string]*service.Post
}

func newMapCache() *MapCache {
	return &MapCache{
		cache: make(map[string]*service.Post),
	}
}

func (m *MapCache) GetFromCache(id string) (*service.Post, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	post, ok := m.cache[id]
	return post, ok
}

func (m *MapCache) AddToCache(id string, post *service.Post) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.cache[id] = post
}

func (m *MapCache) RemoveFromCache(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.cache, id)
}
