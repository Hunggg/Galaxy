package inmemory

import (
	"sync"
	"time"
)

type SimpleStore []*SimpleStoreShard

func NewSimpleStore(shards int) SimpleStore {
	ss := make(SimpleStore, shards)
	for i := range ss {
		ss[i] = &SimpleStoreShard{
			store: make(map[string]*CachedItem),
		}
	}
	return ss
}

type SimpleStoreShard struct {
	mu    sync.RWMutex
	store map[string]*CachedItem
}

func (ss SimpleStore) Set(key string, value interface{}, expires time.Duration) error {
	s := ss[getShard(key, len(ss))]
	s.mu.Lock()
	defer s.mu.Unlock()

	var expireTime int64 = MaxInt64
	if expires >= 0 {
		expireTime = time.Now().Add(expires).Unix()
	}
	item := CachedItem{
		key:        key,
		value:      value,
		expireTime: expireTime,
	}
	s.store[key] = &item
	return nil
}

func (ss SimpleStore) Get(key string) (interface{}, bool) {
	s := ss[getShard(key, len(ss))]
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, ok := s.store[key]
	if !ok {
		return nil, false
	}

	if item.expireTime < time.Now().Unix() {
		return nil, false
	}

	return item.value, true

}

func (ss SimpleStore) Delete(key string) error {
	s := ss[getShard(key, len(ss))]
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.store, key)
	return nil
}

func (ss SimpleStore) CleanUp() {
	for _, s := range ss {
		s.mu.Lock()
		defer s.mu.Unlock()

		for key, item := range s.store {
			if item.expireTime < time.Now().Unix() {
				delete(s.store, key)
			}
		}
	}
}
