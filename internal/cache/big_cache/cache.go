package cache

import (
	"encoding/json"
	"github.com/Olegsuus/Core/internal/service"
	"github.com/allegro/bigcache"
	"log"
	"time"
)

type LibCache struct {
	cache *bigcache.BigCache
}

func NewLibCache() (*LibCache, error) {
	config := bigcache.Config{
		Shards:           16,
		LifeWindow:       10 * time.Minute,
		CleanWindow:      5 * time.Minute,
		MaxEntrySize:     500,
		Verbose:          false,
		HardMaxCacheSize: 64,
	}

	bc, err := bigcache.NewBigCache(config)
	if err != nil {
		return nil, err
	}
	return &LibCache{cache: bc}, nil
}

func (l *LibCache) GetFromCache(id string) (*service.Post, bool) {
	data, err := l.cache.Get(id)
	if err != nil {
		return nil, false
	}

	var post service.Post
	if err = json.Unmarshal(data, &post); err != nil {
		return nil, false
	}
	return &post, true
}

func (l *LibCache) AddToCache(id string, post *service.Post) {
	data, err := json.Marshal(post)
	if err != nil {
		return
	}
	if err = l.cache.Set(id, data); err != nil {
		log.Println(err)
		return
	}
}

func (l *LibCache) RemoveFromCache(id string) {
	if err := l.cache.Delete(id); err != nil {
		log.Println(err)
		return
	}
}
