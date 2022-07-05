package inject

import (
	"context"
	"time"

	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/wearables"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/go-redis/redis/v8"
	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/config"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/fetcher"
	"github.com/spf13/viper"
)

const (
	CacheInterval = 10 * time.Minute
	NodeTimeout   = 10 * time.Second
)

type Injector struct {
	network   string
	chainId   uint64
	rpcClient *rpc.Client
	ethClient *ethclient.Client
	wsClient  *ethclient.Client

	cacheHandler cache.Cache
	metronionDb  metronion.PersistentDb
	configDb     config.PersistentDb
	wearablesDb  wearables.PersistentDb
	redisClient  *redis.Client
	fetcherCore  *fetcher.FetcherCore

	gormConnection *gorm.DB
}

func NewInjector() *Injector {
	config := env.GetConfig()
	injector := &Injector{
		network: config.ChainName,
		chainId: config.ChainId,
	}
	return injector
}

func (c *Injector) Run(ctx context.Context) error {
	fetcherCore := c.ProvideFetcherCore()
	// metadataStorage := c.ProvideMetadataStorage()
	// if err := metadataStorage.Run(); err != nil {
	// 	return err
	// }
	go fetcherCore.Run(ctx)
	return nil
}

func (c *Injector) provideRpcClient(nodeEndpoint string) *rpc.Client {
	if c.rpcClient == nil {
		ctx, cancel := context.WithTimeout(context.Background(), NodeTimeout)
		defer cancel()
		var err error

		c.rpcClient, err = rpc.DialContext(ctx, nodeEndpoint)
		checkErr(err)
	}
	return c.rpcClient
}

func (c *Injector) ProvideEthClient() *ethclient.Client {
	if c.ethClient == nil {
		nodeUrl := viper.GetString("node_url")
		c.ethClient = ethclient.NewClient(c.provideRpcClient(nodeUrl))
	}
	return c.ethClient
}

func (c *Injector) ProvideWsClient() *ethclient.Client {
	if c.wsClient == nil {
		wsUrl := viper.GetString("ws_node_url")
		c.wsClient = ethclient.NewClient(c.provideRpcClient(wsUrl))
	}
	return c.wsClient
}

func (c *Injector) ProvideCache() cache.Cache {
	if c.cacheHandler == nil {
		c.cacheHandler = cache.NewInMemoryCache(CacheInterval)
	}
	return c.cacheHandler
}

func (c *Injector) ProvideFetcherCore() *fetcher.FetcherCore {
	if c.fetcherCore == nil {
		ethClient := c.ProvideEthClient()
		cache := c.ProvideCache()
		var err error
		c.fetcherCore, err = fetcher.NewFetcherCore(ethClient, c.ProvideMetronionDb(), c.ProvideConfigDb(), c.ProvideWearableDb(), cache)
		checkErr(err)
	}
	return c.fetcherCore
}

// func (c *Injector) ProvideNodeSync() *datasync.NodeSync {
// 	if c.nodeSync == nil {
// 		wsClient := c.ProvideWsClient()
// 		cache := c.ProvideCache()
// 		var err error
// 		c.nodeSync, err = datasync.NewNodeSync(wsClient, cache)
// 		checkErr(err)
// 	}
// 	return c.nodeSync
// }

func checkErr(err error) {
	if err != nil {
		// Make sure event is sent
		panic(err)
	}
}