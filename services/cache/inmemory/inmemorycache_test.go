package inmemory

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type InMemoryCacheTestSuite struct {
	suite.Suite
	cache *InMemoryCache
}

func write(cache *InMemoryCache, key string, value string) {
	_ = cache.Set(key, value, 3*time.Second)
}

func writeList(cache *InMemoryCache, key string, value int) {
	_ = cache.LPush(key, value, 10*time.Second)
}

// func concurrentWrite(cache *InMemoryCache, key string, increment int) {
// 	cache.Set(key, increment, 10*time.Second)
// }

func (ts *InMemoryCacheTestSuite) SetupSuite() {
	ts.cache = NewInMemoryCache(2 * time.Second)
}

func (ts *InMemoryCacheTestSuite) TestSimpleStore() {
	assert := ts.Assert()
	write(ts.cache, "key", "value")

	value, ok := ts.cache.Get("key")
	assert.Equal(true, ok)
	assert.Equal("value", value)

	<-time.NewTimer(6 * time.Second).C

	_, ok = ts.cache.Get("key")
	assert.Equal(false, ok)
}

func (ts *InMemoryCacheTestSuite) TestListedStore() {
	assert := ts.Assert()

	for i := 0; i < 15; i++ {
		go writeList(ts.cache, "key", i)
	}

	<-time.NewTimer(1 * time.Second).C
	value, ok := ts.cache.LGetAll("key")
	assert.Equal(true, ok)
	assert.Equal(15, len(value))
}

func TestInMemoryCacheTestSuite(t *testing.T) {
	suite.Run(t, new(InMemoryCacheTestSuite))
}
