package inmemory

import (
	"sync"
	"time"
)

type ListedStore []*ListedStoreShard

func NewListedStore(shards int) ListedStore {
	ls := make([]*ListedStoreShard, shards)
	for i := range ls {
		ls[i] = &ListedStoreShard{
			store: make(map[string][]*CachedItem),
		}
	}

	return ls
}

type ListedStoreShard struct {
	mu    sync.RWMutex
	store map[string][]*CachedItem
}

func (ls ListedStore) LPush(key string, value interface{}, expires time.Duration) error {
	l := ls[getShard(key, len(ls))]
	l.mu.Lock()
	defer l.mu.Unlock()

	cachedList, ok := l.store[key]
	if !ok {
		cachedList = make([]*CachedItem, 0)
	}

	var expireTime int64 = MaxInt64
	if expires >= 0 {
		expireTime = time.Now().Add(expires).Unix()
	}

	cachedList = append(cachedList, &CachedItem{
		value:      value,
		expireTime: expireTime,
	})
	l.store[key] = cachedList

	return nil
}

func (ls ListedStore) LGet(key string, index int) (interface{}, bool) {
	l := ls[getShard(key, len(ls))]
	l.mu.RLock()
	defer l.mu.RUnlock()

	cachedList, ok := l.store[key]
	if !ok {
		return nil, false
	}

	if index > len(cachedList) {
		return nil, false
	}

	item := cachedList[index]
	if item.expireTime > time.Now().Unix() {
		return nil, false
	}

	return item, true
}

func (ls ListedStore) LGetAll(key string) ([]interface{}, bool) {
	l := ls[getShard(key, len(ls))]
	l.mu.RLock()
	defer l.mu.RUnlock()

	cachedList, ok := l.store[key]
	if !ok {
		return nil, false
	}

	copyOfItems := make([]interface{}, 0)
	for _, it := range cachedList {
		if it.expireTime < time.Now().Unix() {
			continue
		}

		copyOfItems = append(copyOfItems, it.value)
	}

	return copyOfItems, true
}

func (ls ListedStore) LDelete(key string) error {
	l := ls[getShard(key, len(ls))]
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.store, key)
	return nil
}

func (ls ListedStore) CleanUp() {
	for _, l := range ls {
		l.mu.Lock()
		defer l.mu.Unlock()
		for key, list := range l.store {
			newItems := make([]*CachedItem, 0, len(list))
			for _, item := range list {
				if item.expireTime < time.Now().Unix() {
					continue
				}

				newItems = append(newItems, item)
			}
			l.store[key] = newItems
		}
	}
}
