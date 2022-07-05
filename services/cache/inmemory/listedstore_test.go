package inmemory

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ListedStoreTestSuite struct {
	suite.Suite

	ls ListedStore
}

func (ts *ListedStoreTestSuite) SetupSuite() {
	ts.ls = NewListedStore(1)
}

func (ts *ListedStoreTestSuite) TestExpiration() {
	assert := ts.Assert()
	err := ts.ls.LPush("key", "value", 1*time.Second)
	assert.NoError(err)

	<-time.NewTimer(2 * time.Second).C
	ts.ls.CleanUp()

	value, ok := ts.ls.LGetAll("key")
	assert.Equal(true, ok)
	assert.Equal(0, len(value))
}

func TestListedStoreTestSuite(t *testing.T) {
	suite.Run(t, new(ListedStoreTestSuite))
}
