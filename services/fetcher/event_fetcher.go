package fetcher

import (
	"sort"
	"sync"

	etherCommon "github.com/ethereum/go-ethereum/common"
	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/metrogalaxy-io/metrogalaxy-api/common"
	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"go.uber.org/zap"
)

type EventFetcher struct {
	l           *zap.SugaredLogger
	config      env.Config
	nodeFetcher *NodeFetcher
}

func NewEventFetcher(nodeFetcher *NodeFetcher) (*EventFetcher, error) {
	l := zap.S()
	config := env.GetConfig()

	return &EventFetcher{
		l:           l,
		config:      config,
		nodeFetcher: nodeFetcher,
	}, nil
}

// FetchBatchRawLogs fetch batch raw logs with multiple routines
func (f *EventFetcher) FetchBatchRawLogs(startBlock, endBlock, batchSize int64, listTopics []etherCommon.Hash, listContracts []etherCommon.Address, numOfRoutines int64) ([]etherTypes.Log, error) {
	wg := &sync.WaitGroup{}
	blockRange := (endBlock - startBlock) / numOfRoutines

	wg.Add(int(numOfRoutines))
	logCh := make(chan []etherTypes.Log)
	errCh := make(chan error)
	doneCh := make(chan bool)

	go func() {
		wg.Wait()
		doneCh <- true
	}()

	// spawn routines
	for i := int64(0); i < numOfRoutines; i++ {
		st := startBlock + i*blockRange
		end := common.MinInt64(st+blockRange-1, endBlock)
		if i == numOfRoutines {
			end = endBlock
		}
		if st > end {
			continue
		}

		go f.fetchRawLogsRoutine(st, end, batchSize, listTopics, listContracts, wg, logCh, errCh)
	}

	logArray := make([]etherTypes.Log, 0)
	counter := int64(0)
L1:
	for {
		select {
		case err := <-errCh:
			f.l.Errorw("err fetch raw logs", "error", err)
			return nil, err
		case rawLogs := <-logCh:
			logArray = append(logArray, rawLogs...)
			counter += 1
		case <-doneCh:
			if counter != numOfRoutines {
				continue
			}
			break L1
		}
	}

	// sort ASC log array
	sort.SliceStable(logArray, func(i, j int) bool {
		return logArray[i].BlockNumber <= logArray[j].BlockNumber
	})

	// f.l.Debugw("fetch raw event logs done", "length", len(logArray))

	return logArray, nil
}

func (f *EventFetcher) fetchRawLogsRoutine(fromBlock, toBlock, batchSize int64, listTopics []etherCommon.Hash, listContracts []etherCommon.Address, wg *sync.WaitGroup, logCh chan []etherTypes.Log, errCh chan error) {
	defer wg.Done()

	// f.l.Debugw("start fetch raw logs routine", "from_block", fromBlock, "to_block", toBlock, "batch_size", batchSize)
	rawLogs, err := f.fetchRawLogs(fromBlock, toBlock, batchSize, listTopics, listContracts)
	if err != nil {
		errCh <- err
		return
	}
	logCh <- rawLogs
}

// FetchRawLogs fetchRawLogs raw logs from all topics of all contracts
func (f *EventFetcher) fetchRawLogs(fromBlock, toBlock, batchSize int64, listTopics []etherCommon.Hash, listContracts []etherCommon.Address) ([]etherTypes.Log, error) {
	topics := [][]etherCommon.Hash{}
	topics = append(topics, listTopics)

	rawLogs, err := f.nodeFetcher.FetchEventsByBatch(fromBlock, toBlock, batchSize, listContracts, topics, func(logs []etherTypes.Log, from, to int64) error { return nil })
	if err != nil {
		f.l.Errorw("error fetch raw logs", "error", err)
		return nil, err
	}
	return rawLogs, nil
}
