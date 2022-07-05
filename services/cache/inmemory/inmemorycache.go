package inmemory

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"time"
)

var (
	MaxUint64     = ^uint64(0)
	MaxInt64      = int64(MaxUint64 >> 1)
	NumberOfShard = 256
)

type InMemoryCache struct {
	simpleStore SimpleStore
	listedStore ListedStore
}

type CachedItem struct {
	key        string
	value      interface{}
	expireTime int64
}

func NewInMemoryCache(cleanInterval time.Duration) *InMemoryCache {
	cache := &InMemoryCache{
		simpleStore: NewSimpleStore(NumberOfShard),
		listedStore: NewListedStore(NumberOfShard),
	}

	go func() {
		ticker := time.NewTicker(cleanInterval)
		for {
			cache.simpleStore.CleanUp()
			cache.listedStore.CleanUp()
			<-ticker.C
		}
	}()
	return cache
}

func (c *InMemoryCache) Set(key string, value interface{}, expires time.Duration) error {
	return c.simpleStore.Set(key, value, expires)
}

func (c *InMemoryCache) Get(key string) (interface{}, bool) {
	return c.simpleStore.Get(key)
}

func (c *InMemoryCache) Delete(key string) error {
	return c.simpleStore.Delete(key)
}

func (c *InMemoryCache) LPush(key string, value interface{}, expires time.Duration) error {
	return c.listedStore.LPush(key, value, expires)
}

func (c *InMemoryCache) LGet(key string, index int) (interface{}, bool) {
	return c.listedStore.LGet(key, index)
}

func (c *InMemoryCache) LGetAll(key string) ([]interface{}, bool) {
	return c.listedStore.LGetAll(key)
}

func (c *InMemoryCache) LDelete(key string) error {
	return c.listedStore.LDelete(key)
}

func getShard(key string, numOfShard int) int {
	// Hash
	hasher := sha1.New()
	_, _ = hasher.Write([]byte(key))
	hash := hasher.Sum(nil)

	// Get the first 8 bytes and convert it to uint32
	buffer := bytes.NewReader(hash[0:4])
	var result uint32
	_ = binary.Read(buffer, binary.BigEndian, &result)

	// Mod it to the number of shard
	return int(result) % numOfShard
}
