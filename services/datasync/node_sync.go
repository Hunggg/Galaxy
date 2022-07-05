package datasync

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache"
	"go.uber.org/zap"
)

const (
	NodeBlockTime = 2 * time.Second
)

type NodeSync struct {
	l        *zap.SugaredLogger
	wsClient *ethclient.Client
	cache    cache.Cache
}

func NewNodeSync(wsClient *ethclient.Client, cache cache.Cache) (*NodeSync, error) {
	return &NodeSync{
		l:        zap.S(),
		wsClient: wsClient,
		cache:    cache,
	}, nil
}

func (n *NodeSync) Run() {
	errCh := make(chan error)
	go n.SubscribeNewHead(errCh)

	err := <-errCh
	n.l.Panic(err)
}

func (n *NodeSync) SubscribeNewHead(errCh chan<- error) {
	sink := make(chan *types.Header)
	sub, err := n.wsClient.SubscribeNewHead(context.Background(), sink)
	if err != nil {
		n.l.Errorw("subscribe new head error", "error", err)
		errCh <- err
		return
	}

	for {
		select {
		case err := <-sub.Err():
			errCh <- err
			return
		case head := <-sink:
			n.l.Debugw("new block number", "block_number", head.Number.Uint64())
		}
	}
}
