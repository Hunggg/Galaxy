package common

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	etherCommon "github.com/ethereum/go-ethereum/common"
	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache"
)

const (
	BlockTimeExpired          = 2 * time.Second
	GetBlockHeaderCachePrefix = "get_block_header"
)

func DialRpcNode(nodeEndpoint string) (*rpc.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), NodeTimeout)
	defer cancel()
	return rpc.DialContext(ctx, nodeEndpoint)
}

func CreateEthClient(nodeEndpoint string) (*ethclient.Client, error) {
	rpcClient, err := DialRpcNode(nodeEndpoint)
	if err != nil {
		return nil, err
	}
	return ethclient.NewClient(rpcClient), nil
}

// GetBlockHeader get block header from cache if exists, else fetch from node
func GetBlockHeader(cache cache.Cache, client *ethclient.Client, blockNumber *big.Int) (*etherTypes.Header, error) {
	key := fmt.Sprintf("%s_%s", GetBlockHeaderCachePrefix, blockNumber.String())
	v, ok := cache.Get(key)

	if ok {
		return v.(*etherTypes.Header), nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), NodeTimeout)
	defer cancel()
	header, err := client.HeaderByNumber(ctx, blockNumber)

	if err != nil {
		return nil, err
	} else {
		_ = cache.Set(key, header, BlockTimeExpired)
		return header, err
	}
}

// FetchEventsByBatch fetch event logs
func FetchEventsByBatch(ethClient *ethclient.Client,
	fromBlock, toBlock, batchSize int64,
	contracts []etherCommon.Address,
	topics [][]etherCommon.Hash,
	logParser func([]etherTypes.Log, int64, int64) error) ([]etherTypes.Log, error) {
	result := []etherTypes.Log{}
	for fromBlock <= toBlock {
		to := MinInt64(fromBlock+batchSize-1, toBlock)

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(fromBlock),
			ToBlock:   big.NewInt(to),
			Addresses: contracts,
			Topics:    topics,
		}

		logs, err := ethClient.FilterLogs(context.Background(), query)
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
