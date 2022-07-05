package config

import (
	"testing"

	"github.com/metrogalaxy-io/metrogalaxy-api/libs/cockroachdb"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type CockroachDbTestSuite struct {
	suite.Suite
	client *CockroachDB
}

func TestCockroachDbTestSuite(t *testing.T) {
	suite.Run(t, new(CockroachDbTestSuite))
}

func (ts *CockroachDbTestSuite) SetupSuite() {
	assert := ts.Assert()
	viper.Set("METROGALAXY_API_COCKROACH_DB_URI", "postgresql://root@localhost:26257/metrogalaxy_api_test?sslmode=disable")
	db, err := cockroachdb.NewCockroachDbConnection()
	assert.NoError(err)
	assert.NotNil(db)
	client, err := NewCockroachDB(db)
	assert.NoError(err)
	assert.NotNil(client)

	ts.client = client
}

// func (ts *CockroachDbTestSuite)
