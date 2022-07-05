// +build !skipTest

package cockroachdb

import (
	"strings"
	"testing"

	"github.com/metrogalaxy-io/metrogalaxy-api/libs/cockroachdb"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"
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

func (ts *CockroachDbTestSuite) TearDownAllSuite() {
	_ = ts.client.db.Migrator().DropTable(&metronion.MetronionVersionConfig{})
	_ = ts.client.db.Migrator().DropTable(&metronion.MetronionPrice{})
	_ = ts.client.db.Migrator().DropTable(&metronion.MetronionListing{})
	_ = ts.client.db.Migrator().DropTable(&metronion.MetronionOffers{})
}

func (ts *CockroachDbTestSuite) TestSaveBatchMetronionMetadata() {
	assert := ts.Assert()

	data := make([]metronion.MetronionMetadata, 5)
	for i := 0; i < 5; i++ {
		data[i] = metronion.MetronionMetadata{
			Id: uint64(i),
		}
	}

	err := ts.client.SaveBatchMetronionMetadata(data)
	assert.NoError(err)
}

func (ts *CockroachDbTestSuite) TestSaveMetronionVersionConfig() {
	assert := ts.Assert()

	data := metronion.MetronionVersionConfig{
		VersionID:     0,
		StartingIndex: 0,
		Provenance:    "0x123456",
		MaxSupply:     10000,
		CurrentSupply: 0,
	}
	err := ts.client.SaveMetronionVersionConfig(data)
	assert.NoError(err)

	result, err := ts.client.GetMetronionVersionConfig(0)
	assert.NoError(err)
	assert.NotNil(result)
	assert.Equal(data.MaxSupply, result.MaxSupply)
}

func (ts *CockroachDbTestSuite) TestUpdateMetronionPrice() {
	assert := ts.Assert()
	var (
		metronionId  uint64 = 0
		priceType           = metronion.MetronionPriceTypeCurrentPrice
		currentPrice        = 10.123
	)

	// insert current price
	err := ts.client.UpdateMetronionPrice(metronionId, priceType, currentPrice)
	assert.NoError(err)

	res, err := ts.client.GetMetronionPrice(metronionId)
	assert.NoError(err)
	assert.NotNil(res)
	assert.Equal(metronionId, res.MetronionID)
	assert.Equal(currentPrice, res.CurrentPrice)

	// upsert
	var newCurrentPrice = 20.456
	err = ts.client.UpdateMetronionPrice(metronionId, priceType, newCurrentPrice)
	assert.NoError(err)

	res, err = ts.client.GetMetronionPrice(metronionId)
	assert.NoError(err)
	assert.NotNil(res)
	assert.Equal(metronionId, res.MetronionID)
	assert.Equal(newCurrentPrice, res.CurrentPrice)

	// insert last price
	var lastPrice = 5.123
	err = ts.client.UpdateMetronionPrice(metronionId, metronion.MetronionPriceTypeLastPrice, lastPrice)
	assert.NoError(err)

	res, err = ts.client.GetMetronionPrice(metronionId)
	assert.NoError(err)
	assert.NotNil(res)
	assert.Equal(metronionId, res.MetronionID)
	assert.Equal(newCurrentPrice, res.CurrentPrice)
	assert.Equal(lastPrice, res.LastPrice)
}

func (ts *CockroachDbTestSuite) TestUpdateMetronionListing() {
	assert := ts.Assert()

	var (
		metronionId uint64 = 0
	)

	data := metronion.MetronionListing{
		MetronionID: metronionId,
		FromAccount: "0xabcAdASDASdef",
		Price:       1.2345,
		Timestamp:   1652108721,
		BlockNumber: 123456,
		TxHash:      "0x789def",
	}

	err := ts.client.UpdateMetronionListing(data, true)
	assert.NoError(err)

	res, err := ts.client.GetMetronionListing(metronionId)
	assert.NoError(err)
	assert.NotNil(res)
	assert.Equal(data.MetronionID, res.MetronionID)
	assert.Equal(strings.ToLower(data.FromAccount), res.FromAccount)
	assert.Equal(data.Price, res.Price)
	assert.Equal(data.Timestamp, res.Timestamp)
	assert.Equal(data.BlockNumber, res.BlockNumber)
	assert.Equal(data.TxHash, res.TxHash)

	err = ts.client.UpdateMetronionListing(data, true)
	assert.Error(err)

	err = ts.client.UpdateMetronionListing(data, false)
	assert.NoError(err)

	res, err = ts.client.GetMetronionListing(metronionId)
	assert.Error(err)
}

func (ts *CockroachDbTestSuite) TestUpdateMetronionOffers() {
	assert := ts.Assert()

	var (
		metronionId uint64 = 0
	)

	data := metronion.MetronionOffers{
		MetronionID: metronionId,
		FromAccount: "0xabcAdASDASdef",
		Price:       1.2345,
		Timestamp:   1652108721,
		BlockNumber: 123456,
		TxHash:      "0x789def",
	}

	err := ts.client.UpdateMetronionOffers(data, true)
	assert.NoError(err)

	resArr, err := ts.client.GetMetronionOffers(metronionId, "")
	assert.NoError(err)
	assert.NotNil(resArr)
	assert.Equal(1, len(resArr))
	res := resArr[0]

	assert.Equal(data.MetronionID, res.MetronionID)
	assert.Equal(strings.ToLower(data.FromAccount), res.FromAccount)
	assert.Equal(data.Price, res.Price)
	assert.Equal(data.Timestamp, res.Timestamp)
	assert.Equal(data.BlockNumber, res.BlockNumber)
	assert.Equal(data.TxHash, res.TxHash)

	err = ts.client.UpdateMetronionOffers(data, true)
	assert.Error(err)

	// insert other record
	newData := metronion.MetronionOffers{
		MetronionID: metronionId,
		FromAccount: "0xaASDasbcAdASDASdef",
		Price:       1.2345,
		Timestamp:   1652108721,
		BlockNumber: 123456,
		TxHash:      "0x789def",
	}
	err = ts.client.UpdateMetronionOffers(newData, true)
	assert.NoError(err)

	resArr, err = ts.client.GetMetronionOffers(metronionId, "")
	assert.NoError(err)
	assert.NotNil(resArr)
	assert.Equal(2, len(resArr))
	res = resArr[1]

	assert.Equal(newData.MetronionID, res.MetronionID)
	assert.Equal(strings.ToLower(newData.FromAccount), res.FromAccount)
	assert.Equal(newData.Price, res.Price)
	assert.Equal(newData.Timestamp, res.Timestamp)
	assert.Equal(newData.BlockNumber, res.BlockNumber)
	assert.Equal(newData.TxHash, res.TxHash)

	// teardown
	ts.client.db.Delete(&metronion.MetronionOffers{})
}
