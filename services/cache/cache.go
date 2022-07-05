package cache

import (
	"time"

	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache/inmemory"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache/rediscache"
)

type Cache interface {
	// Set an item to cache. expires = 0 means default expiration, -1 means infinite
	Set(key string, value interface{}, expires time.Duration) error
	Get(key string) (interface{}, bool)
	Delete(key string) error

	// Mimic redis
	LPush(key string, value interface{}, expires time.Duration) error
	LGet(key string, index int) (interface{}, bool)
	LGetAll(key string) ([]interface{}, bool)
	LDelete(key string) error
}

func NewInMemoryCache(cleanUpInterval time.Duration) *inmemory.InMemoryCache {
	return inmemory.NewInMemoryCache(cleanUpInterval)
}

func NewRedisCache() *rediscache.CacheClient {
	return rediscache.NewCacheClient()
}
