package cockroachdb

import (
	"fmt"
	"testing"

	"github.com/metrogalaxy-io/metrogalaxy-api/libs/cockroachdb"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/wearables"
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
	viper.Set("METROGALAXY_API_COCKROACH_DB_URI", "postgresql://root@localhost:26257/metrogalaxy_api?sslmode=disable")
	db, err := cockroachdb.NewCockroachDbConnection()
	assert.NoError(err)
	assert.NotNil(db)
	client, err := NewCockroachDB(db)
	assert.NoError(err)
	assert.NotNil(client)

	ts.client = client
}

func (ts *CockroachDbTestSuite) TestGetWearablesListingsById() {
	assert := ts.Assert()
	var txhash string = "0x3a06879873bbcdd3dd020d737f154c11c572dd776e7b9719fe83ed571a1bebad"
	var id uint64 = 1

	res, err := ts.client.GetWearablesListing(id)
	assert.NoError(err)
	assert.NotNil(res)
	assert.Equal(txhash, string(res.TxHash))	
}

func (ts *CockroachDbTestSuite) TestGetWearablesOfferssById() {
	assert := ts.Assert()
	var txhash string = "0xb80fafe1363fb7d9f4c54d9a98e787fc1d668a01b69c7abf51f20edd0a51146d"
	var id uint64 = 12

	res, err := ts.client.GetWearablesOffer(id)
	assert.NoError(err)
	assert.NotNil(res)
	assert.Equal(txhash, string(res.TxHash))	
}

func (ts *CockroachDbTestSuite) TestGetWearablesOnChainById() {
	assert := ts.Assert()
	target := wearables.WearablesOnChainData  {
		Id: 1,
		Name: "Red Hair",
		MaxSupply: 1000,
		CurrentSupply: 25,
		Type: "hair",
		Rarity: "common",
		CreatedAtTimestamp: 1653973171,
		UpdatedAtTimestamp: 1654228349,
		CreatedAtBlock: 10239798,
		UpdatedAtBlock: 10341590,
	}

	res, err := ts.client.getWearablesOnChainDataById(target.Id)
	assert.NoError(err)
	assert.NotNil(res)
	assert.Equal(target.Id, res.Id)	
	assert.Equal(target.Name, res.Name)	
	assert.Equal(target.MaxSupply, res.MaxSupply)	
	assert.Equal(target.CurrentSupply, res.CurrentSupply)	
	assert.Equal(target.Type, res.Type)	
	assert.Equal(target.Rarity, res.Rarity)	
	assert.Equal(target.CreatedAtTimestamp, res.CreatedAtTimestamp)	
	assert.Equal(target.UpdatedAtTimestamp, res.UpdatedAtTimestamp)	
	assert.Equal(target.UpdatedAtBlock, res.UpdatedAtBlock)	
	assert.Equal(target.CreatedAtBlock, res.CreatedAtBlock)	
}

func (ts *CockroachDbTestSuite) TestGetListWearablesOnChainData() {
	assert := ts.Assert()
	var offset = 1
	var limit = 5

	res, err := ts.client.getListWearablesOnchainData(offset, limit)
	assert.NoError(err)
	assert.NotNil(res)
	assert.Equal(uint64(offset), res[0].Id)
	assert.Equal(limit, len(res))
}

func (ts *CockroachDbTestSuite) TestGetListWearablesListing() {
	assert := ts.Assert()
	var offset = 0
	var limit = 2

	res, count, err := ts.client.GetListWearablesListing(offset, limit)
	assert.NoError(err)
	assert.NotNil(res)
	fmt.Println("count: ", count)
	assert.NotEqual(count, uint64(0))
	assert.Equal(limit, len(res))
}

func (ts *CockroachDbTestSuite) TestGetListWearablesOffers() {
	assert := ts.Assert()
	var offset = 0
	var limit = 1

	res, count, err := ts.client.GetListWearablesOffer(offset, limit)
	assert.NoError(err)
	assert.NotNil(res)
	fmt.Println("count: ", count)
	assert.NotEqual(count, uint64(0))
	assert.Equal(limit, len(res))
}

func (ts *CockroachDbTestSuite) TestGetListWearablesActivities() {
	assert := ts.Assert()
	var offset = 0
	var limit = 2

	res, count, err := ts.client.GetListWearablesActivity(offset, limit)
	assert.NoError(err)
	assert.NotNil(res)
	fmt.Println("count :", count)
	assert.NotEqual(count, uint64(0))
	assert.Equal(limit, len(res))
}

func (ts *CockroachDbTestSuite) TestGetWearablesInfo() {
	assert := ts.Assert()
	var id = 17
	res, err := ts.client.GetWearablesInformation(uint64(id))
	assert.NoError(err)
	assert.NotNil(res)
	fmt.Println(res.CurrentSupply)
}

func (ts *CockroachDbTestSuite) TestGetListWearablesInfor() {
	assert := ts.Assert()
	var offset = 0
	var limit = 2
	filter := wearables.WearablesFilterParams {
		Rarity: "common",
		Sort: wearables.WearablesSortingLowestCurrentSupply,
		Gender: nil,
		Type: nil,
	}

	res, count, err := ts.client.GetListWearablesInformation(offset, limit, filter)
	fmt.Println(count)
	assert.NoError(err)
	assert.NotNil(res)
}
