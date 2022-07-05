package fetcher

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	etherCommon "github.com/ethereum/go-ethereum/common"
	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metrogalaxy-io/metrogalaxy-api/common"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/limiter"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache"
	"go.uber.org/zap"
)

const (
	DEFAULT_RATE_LIMIT         = 5 //rps
	DEFAULT_BURST              = 10
	DEFAULT_RATE_LIMIT_TIMEOUT = 30 * time.Second

	BlockTimeExpired          = 2 * time.Second
	GetBlockHeaderCachePrefix = "get_block_header"
)

type NodeFetcher struct {
	l           *zap.SugaredLogger
	cache       cache.Cache
	rateLimiter *limiter.RateLimiter
	ethClient   *ethclient.Client
}

func NewNodeFetcher(ethClient *ethclient.Client, cache cache.Cache) *NodeFetcher {
	return &NodeFetcher{
		l:           zap.S(),
		cache:       cache,
		ethClient:   ethClient,
		rateLimiter: limiter.NewRateLimiter(DEFAULT_RATE_LIMIT, DEFAULT_BURST),
	}
}

func (f *NodeFetcher) GetBlockHeader(blockNumber *big.Int) (*etherTypes.Header, error) {
	key := fmt.Sprintf("%s_%s", GetBlockHeaderCachePrefix, blockNumber.String())
	v, ok := f.cache.Get(key)

	if ok {
		return v.(*etherTypes.Header), nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), common.NodeTimeout)
	defer cancel()
	if err := f.rateLimiter.WaitN(DEFAULT_RATE_LIMIT_TIMEOUT, 1); err != nil {
		return nil, err
	}
	header, err := f.ethClient.HeaderByNumber(ctx, blockNumber)

	if err != nil {
		return nil, err
	} else {
		_ = f.cache.Set(key, header, BlockTimeExpired)
		return header, err
	}
}

// FetchEventsByBatch fetch raw logs in range of block number with batch size
func (f *NodeFetcher) FetchEventsByBatch(
	fromBlock, toBlock, batchSize int64,
	contracts []etherCommon.Address,
	topics [][]etherCommon.Hash,
	logParser func([]etherTypes.Log, int64, int64) error) ([]etherTypes.Log, error) {
	f.l.Debugw("fetch event logs by batch", "from_block", fromBlock, "to_block", toBlock, "batch_size", batchSize)
	result := []etherTypes.Log{}
	for fromBlock <= toBlock {
		to := common.MinInt64(fromBlock+batchSize-1, toBlock)

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock),
			ToBlock:   big.NewInt(to),
			Addresses: contracts,
			Topics:    topics,
		}

		if err := f.rateLimiter.WaitN(DEFAULT_RATE_LIMIT_TIMEOUT, 1); err != nil {
			return nil, err
		}

		logs, err := f.ethClient.FilterLogs(context.Background(), query)
		if err != nil {
			return result, err
		}

		err = logParser(logs, fromBlock, to)
		if err != nil {
			return result, err
		}

		result = append(result, logs...)

		fromBlock = to + 1
	}

	return result, nil
}
