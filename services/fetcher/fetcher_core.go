package fetcher

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/wearables"

	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metrogalaxy-io/metrogalaxy-api/common"
	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/eventbus"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/config"
	"go.uber.org/zap"
)

type FetcherClientName int

const (
	DEFAULT_NUMBER_OF_ROUTINE = int64(5)
	BLOCK_RANGE               = 2000
)

const (
	EventBusBuffered = 500
	BlockGap         = 6

	MetronionNFTFetcherClient FetcherClientName = iota
	MetronionSaleFetcherClient
	MetronionAccessoriesFetcherClient
)

type EventConsumer interface {
}

type FetcherClient interface {
	Run(ctx context.Context, errCh chan<- error)
	FetcherClientName() FetcherClientName
	FetchRawLogs(startBlock, endBlock, batchSize, numOfRoutines int64) ([]etherTypes.Log, error)
}

type EventHandler interface {
	Run(ctx context.Context, errCh chan<- error)
}

type FetcherCore struct {
	l           *zap.SugaredLogger
	config      env.Config
	ethClient   *ethclient.Client
	nodeFetcher *NodeFetcher

	metronionDb        metronion.PersistentDb
	configDb           config.PersistentDb
	wearableDb         wearables.PersistentDb
	cache              cache.Cache
	subscriptionKeeper *eventbus.SubscriptionKeeper

	clients       []FetcherClient
	eventHandlers []EventHandler
}

func NewFetcherCore(ethClient *ethclient.Client, metronionDb metronion.PersistentDb, configDb config.PersistentDb, wearableDb wearables.PersistentDb, cache cache.Cache) (*FetcherCore, error) {
	l := zap.S()
	config := env.GetConfig()

	subscriptionKeeper := eventbus.NewSubscriptionKeeper(EventBusBuffered)

	clients := make([]FetcherClient, 0)

	nodeFetcher := NewNodeFetcher(ethClient, cache)

	eventFetcher, err := NewEventFetcher(nodeFetcher)
	if err != nil {
		l.Errorw("error create event fetcher", "error", err)
		return nil, err
	}

	metronionNFTFetcher, err := NewMetronionNFTFetcher(ethClient, eventFetcher, metronionDb, cache, subscriptionKeeper)
	if err != nil {
		l.Errorw("error create MetronionNFT fetcher", "error", err)
		return nil, err
	}

	marketplaceFetcher, err := NewMarketplaceFetcher(ethClient, eventFetcher, cache, subscriptionKeeper)
	if err != nil {
		l.Errorw("error create Marketplace fetcher", "error", err)
		return nil, err
	}

	wearableFetcher, err := NewWearablesFetcher(ethClient, eventFetcher, wearableDb, cache, subscriptionKeeper)
	if err != nil {
		l.Errorw("error create Wearables fetcher", "error", err)
		return nil, err
	}

	clients = append(clients, metronionNFTFetcher, marketplaceFetcher, wearableFetcher)

	metronionHandler := NewMetronionHandler(metronionDb, subscriptionKeeper)
	metronionActivityHandler := NewMetronionActivityHandler(metronionDb, subscriptionKeeper)
	metronionOffersHandler := NewMetronionOffersHandler(metronionDb, subscriptionKeeper)
	metronionListingHandler := NewMetronionListingHandler(metronionDb, subscriptionKeeper)
	wearableHandler := NewWearableHandler(wearableDb, subscriptionKeeper)

	eventHandlers := []EventHandler{metronionHandler, metronionActivityHandler, metronionOffersHandler, metronionListingHandler, wearableHandler}

	return &FetcherCore{
		l:                  l,
		config:             config,
		ethClient:          ethClient,
		metronionDb:        metronionDb,
		configDb:           configDb,
		cache:              cache,
		subscriptionKeeper: subscriptionKeeper,
		clients:            clients,
		eventHandlers:      eventHandlers,
		nodeFetcher:        nodeFetcher,
	}, nil
}

func (c *FetcherCore) Run(ctx context.Context) {
	errCh := make(chan error)
	// subscribe to events
	for _, handler := range c.eventHandlers {
		handler.Run(ctx, errCh)
	}

	// subscribe to raw logs
	for _, client := range c.clients {
		client.Run(ctx, errCh)
	}

	go c.RunFetch(ctx, errCh)

	for {
		err := <-errCh
		c.l.Errorw("fetcher core error", "error", err)
		c.l.Panic(err)
	}
}

// RunFetch run continuously to fetch raw logs and dispatch to multiple handlers
func (c *FetcherCore) RunFetch(ctx context.Context, errCh chan error) {
	var startFetchingBlock = c.config.MetronionNFTContract.CreatedAtBlock
	dbConfig, err := c.configDb.GetConfig()
	if err != nil {
		c.l.Errorw("error get database config", "error", err)
		errCh <- err
		return
	}

	startFetchingBlock = FindMaxBlocks(startFetchingBlock, dbConfig.LastFetchedBlock)

	c.l.Infow("run fetcher core", "start_block", startFetchingBlock)

	var fromBlock = int64(startFetchingBlock)
	ticker := time.NewTicker(common.BlockTime)
	defer ticker.Stop()

	numberOfRoutine := DEFAULT_NUMBER_OF_ROUTINE

	for {
		select {
		case <-ctx.Done():
			c.l.Infow("stopping fetcher core")
			return
		default:
			blockHeader, err := c.nodeFetcher.GetBlockHeader(nil)
			if err != nil {
				c.l.Errorw("get block header error", "error", err)
				<-ticker.C
				continue
			}

			toBlock := blockHeader.Number.Int64()
			if toBlock-fromBlock < BlockGap {
				<-ticker.C
				// c.l.Debugw("skip", "from_block", fromBlock, "to_block", toBlock)
				continue
			}
			toBlock = common.MinInt64(fromBlock+BLOCK_RANGE, toBlock)

			// each fetcher client fetches raw logs and we combine them into an array of raw logs
			rawLogs, err := c.FetchRawLogs(fromBlock, toBlock, common.BlockRange, numberOfRoutine)
			if err != nil {
				<-ticker.C
				continue
			}

			c.l.Debugw("raw logs", "length", len(rawLogs))

			// publish raw logs
			if err := c.PublishRawLogs(rawLogs); err != nil {
				<-ticker.C
				continue
			}

			// update last fetched block number
			c.l.Infow("syncing block successfully", "block_number", toBlock)
			if err := c.configDb.UpdateConfig(map[config.ConfigField]interface{}{
				config.ConfigFieldLastFetchedBlock: uint64(toBlock),
			}); err != nil {
				c.l.Errorw("error save last fetched block", "error", err)
				<-ticker.C
				continue
			}

			fromBlock = toBlock + 1
			// for interval fetch, we only need to use 2 routines
			if numberOfRoutine == int64(10) {
				numberOfRoutine = int64(2)
			}
		}
	}
}

// FetchRawLogs Fetch raw logs from each clients, and combine those result into sorted array
func (c *FetcherCore) FetchRawLogs(startBlock, endBlock, batchSize, numOfRoutines int64) ([]etherTypes.Log, error) {
	wg := sync.WaitGroup{}
	rawLogCh := make(chan etherTypes.Log)
	result := make([]etherTypes.Log, 0)
	doneCh := make(chan struct{})
	errCh := make(chan error)

	for i := range c.clients {
		wg.Add(1)
		go c.fetchRoutine(c.clients[i], startBlock, endBlock, batchSize, numOfRoutines, rawLogCh, &wg, errCh)
	}

	go func() {
		wg.Wait()
		doneCh <- struct{}{}
	}()

	for {
		select {
		case <-doneCh:
			sort.SliceStable(result, func(i, j int) bool {
				return result[i].BlockNumber <= result[j].BlockNumber
			})
			return result, nil
		case err := <-errCh:
			c.l.Errorw("error fetch log", "error", err)
			return nil, err
		case rawLog := <-rawLogCh:
			result = append(result, rawLog)
		}
	}
}

func (c *FetcherCore) fetchRoutine(client FetcherClient, startBlock, endBlock, batchSize, numOfRoutines int64, rawLogCh chan etherTypes.Log, wg *sync.WaitGroup, errCh chan error) {
	defer wg.Done()
	rawLogs, err := client.FetchRawLogs(startBlock, endBlock, batchSize, numOfRoutines)
	if err != nil {
		errCh <- err
		return
	}
	for i := range rawLogs {
		rawLogCh <- rawLogs[i]
	}
}

func (c *FetcherCore) PublishRawLogs(rawLogs []etherTypes.Log) error {
	for _, item := range rawLogs {
		c.subscriptionKeeper.Publish(RawLogSubscriptionTopic, item)
	}
	return nil
}
