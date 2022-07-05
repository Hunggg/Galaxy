package fetcher

import (
	"context"
	"fmt"
	"math/big"

	etherCommon "github.com/ethereum/go-ethereum/common"
	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metrogalaxy-io/metrogalaxy-api/common"
	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/contracts"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/eventbus"
	"github.com/metrogalaxy-io/metrogalaxy-api/services"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/wearables"
	"go.uber.org/zap"
)

type WearablesFetcher struct {
	l         *zap.SugaredLogger
	config    env.Config
	ethClient *ethclient.Client

	metronionAccessoriesAddress etherCommon.Address
	metronionAccessories        *contracts.MetronionAccessories
	topics                      []etherCommon.Hash
	listContracts               []etherCommon.Address
	persistentDb                wearables.PersistentDb
	cache                       cache.Cache
	eventFetcher                *EventFetcher
	subscriptionKeeper          *eventbus.SubscriptionKeeper
}

func NewWearablesFetcher(ethClient *ethclient.Client, eventFetcher *EventFetcher, wearableDb wearables.PersistentDb, cache cache.Cache, subscriptionKeeper *eventbus.SubscriptionKeeper) (*WearablesFetcher, error) {
	l := zap.S()
	config := env.GetConfig()
	metronionAccessories, err := contracts.NewMetronionAccessories(etherCommon.HexToAddress(config.MetronionAccessoriesContract.Address), ethClient)
	if err != nil {
		l.Errorw("initialize metronion accessories contract client error", "error", err)
		return nil, err
	}

	topics := []etherCommon.Hash{
		contracts.MetronionAccessories_AccessoryCreatedTopic,
		contracts.MetronionAccessories_TransferSingleTopic,
		contracts.MetronionAccessories_TransferBatchTopic,
		contracts.MetronionAccessories_AccessoryMintTopic,
		contracts.MetronionAccessories_AccessoryBurntTopic,
		contracts.MetronionAccessories_AccessoryStoreTopic,
		contracts.MetronionAccessories_AccessoryReturnTopic,
	}
	listContracts := []etherCommon.Address{
		etherCommon.HexToAddress(config.MetronionAccessoriesContract.Address),
	}
	return &WearablesFetcher{
		l:                           l,
		config:                      config,
		ethClient:                   ethClient,
		metronionAccessoriesAddress: etherCommon.HexToAddress(config.MetronionAccessoriesContract.Address),
		metronionAccessories:        metronionAccessories,
		topics:                      topics,
		listContracts:               listContracts,
		persistentDb:                wearableDb,
		cache:                       cache,
		eventFetcher:                eventFetcher,
		subscriptionKeeper:          subscriptionKeeper,
	}, nil
}

func (c *WearablesFetcher) Run(ctx context.Context, errCh chan<- error) {
	go c.SubscribeRawLogs(ctx, errCh)
}

func (c *WearablesFetcher) FetcherClientName() FetcherClientName {
	return MetronionAccessoriesFetcherClient
}

func (c *WearablesFetcher) FetchRawLogs(startBlock, endBlock, batchSize, numOfRoutines int64) ([]etherTypes.Log, error) {
	c.l.Debugw("wearables fetcher fetch raw logs", "from_block", startBlock, "to_block", endBlock, "batch_size", batchSize)
	return c.eventFetcher.FetchBatchRawLogs(startBlock, endBlock, batchSize, c.topics, c.listContracts, numOfRoutines)
}

/**
 * Parsing methods
 */

func (c *WearablesFetcher) SubscribeRawLogs(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := c.subscriptionKeeper.Subscribe(RawLogSubscriptionTopic, internalErrCh)

	c.l.Infow("subscribe raw logs event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			c.l.Debugw("unsubscribe raw logs event", "sub_id", sub.Id)
			c.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			c.l.Errorw("subscribe wearables raw logs event error. Retry subscribing...", "error", err)
			sub = c.subscriptionKeeper.Subscribe(RawLogSubscriptionTopic, internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(etherTypes.Log)
			c.l.Debugw("receive raw log event")
			if !ok {
				c.l.Errorw("error decode raw logs event", "message", message)
				errCh <- fmt.Errorf("error decode raw logs event")
				return
			}
			if err := c.ParseRawLog(value); err != nil {
				c.l.Errorw("error parse raw logs event", "error", err)
				errCh <- fmt.Errorf("error parse raw logs event")
				return
			}
		}
	}
}

// ParseRawLog parse raw logs should be in correct order
func (c *WearablesFetcher) ParseRawLog(rawLog etherTypes.Log) error {
	wearableCreatedEvent, err := c.ParseWearableCreatedEvent(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionAccessories_AccessoryCreatedTopic.Hex(), wearableCreatedEvent)
		c.l.Infow("wearable created event", "wearble_on_chain_id", wearableCreatedEvent.OnChainId, "name", wearableCreatedEvent.Name,
			"created_at_timestamp", wearableCreatedEvent.Timestamp, "block_number", wearableCreatedEvent.BlockNumber)
		return nil
	}
	if err != common.ErrorCannotParseLog {
		c.l.Error(err)
		return err
	}

	wearableMintedEvent, err := c.ParseWearableMintEvent(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionAccessories_AccessoryMintTopic.Hex(), wearableMintedEvent)
		c.l.Infow("wearable mint event", "wearble_on_chain_id", wearableMintedEvent.OnChainId, "amount", wearableMintedEvent.Amount,
			"to", wearableMintedEvent.To,
			"created_at_timestamp", wearableMintedEvent.Timestamp, "block_number", wearableMintedEvent.BlockNumber)
		return nil
	}
	if err != common.ErrorCannotParseLog {
		c.l.Error(err)
		return err
	}

	wearableBurntEvent, err := c.ParseWearableBurntEvent(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionAccessories_AccessoryBurntTopic.Hex(), wearableBurntEvent)
		c.l.Infow("wearable burnt event", "wearble_on_chain_id", wearableBurntEvent.OnChainId, "amount", wearableBurntEvent.Amount,
			"from", wearableBurntEvent.From,
			"created_at_timestamp", wearableBurntEvent.Timestamp, "block_number", wearableBurntEvent.BlockNumber)
		return nil
	}
	if err != common.ErrorCannotParseLog {
		c.l.Error(err)
		return err
	}

	wearbleStoredEvent, err := c.ParseWearableStoredEvent(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionAccessories_AccessoryStoreTopic.Hex(), wearbleStoredEvent)
		c.l.Infow("wearable stored event", "wearble_on_chain_id",
			"from", wearbleStoredEvent.From,
			"created_at_timestamp", wearbleStoredEvent.Timestamp, "block_number", wearbleStoredEvent.BlockNumber)
		return nil
	}
	if err != common.ErrorCannotParseLog {
		c.l.Error(err)
		return err
	}

	wearableReturnedEvent, err := c.ParseWearableStoredEvent(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionAccessories_AccessoryReturnTopic.Hex(), wearableReturnedEvent)
		c.l.Infow("wearable returned event", "wearble_on_chain_id",
			"from", wearableReturnedEvent.From,
			"created_at_timestamp", wearableReturnedEvent.Timestamp, "block_number", wearableReturnedEvent.BlockNumber)
		return nil
	}
	if err != common.ErrorCannotParseLog {
		c.l.Error(err)
		return err
	}

	singleTransferEvent, err := c.ParseWearableTransferEvent(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionAccessories_TransferSingleTopic.Hex(), singleTransferEvent)
		c.l.Infow("wearable single transfer event", "wearble_on_chain_id", singleTransferEvent.OnChainId, "amount", singleTransferEvent.Amount,
			"from", singleTransferEvent.From, "to", singleTransferEvent.To,
			"created_at_timestamp", singleTransferEvent.Timestamp, "block_number", singleTransferEvent.BlockNumber)
		return nil
	}
	if err != common.ErrorCannotParseLog {
		c.l.Error(err)
		return err
	}

	batchTransferEvent, err := c.ParseWearableBatchTransferEvent(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionAccessories_TransferBatchTopic.Hex(), batchTransferEvent)
		c.l.Infow("wearable batch transfer event", "wearble_on_chain_ids", batchTransferEvent.OnChainIds, "amounts", batchTransferEvent.Amounts,
			"from", batchTransferEvent.From, "to", batchTransferEvent.To,
			"created_at_timestamp", batchTransferEvent.Timestamp, "block_number", batchTransferEvent.BlockNumber)
		return nil
	}
	if err != common.ErrorCannotParseLog {
		c.l.Error(err)
		return err
	}
	return nil
}

func (c *WearablesFetcher) ParseWearableCreatedEvent(rawLog etherTypes.Log) (wearables.WearableCreatedEvent, error) {
	wearableCreatedEvent, err := c.metronionAccessories.ParseAccessoryCreated(rawLog)
	if err != nil {
		return wearables.WearableCreatedEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(wearableCreatedEvent.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableCreatedEvent{}, err
	}

	return wearables.WearableCreatedEvent{
		OnChainId:   wearableCreatedEvent.Id.Uint64(),
		Name:        wearableCreatedEvent.Name,
		MaxSupply:   wearableCreatedEvent.MaxSupply.Uint64(),
		Type:        ParseWearablesType(wearableCreatedEvent.AccessoriesType.Uint64()),
		Rarity:      ParseRarity(wearableCreatedEvent.Rarity),
		BlockNumber: blockHeader.Number.Uint64(),
		Timestamp:   blockHeader.Time,
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *WearablesFetcher) ParseWearableMintEvent(rawLog etherTypes.Log) (wearables.WearableMintedEvent, error) {
	event, err := c.metronionAccessories.ParseAccessoryMint(rawLog)
	if err != nil {
		return wearables.WearableMintedEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableMintedEvent{}, err
	}

	return wearables.WearableMintedEvent{
		OnChainId:   event.Id.Uint64(),
		Amount:      event.Amount.Uint64(),
		To:          event.To.Hex(),
		BlockNumber: blockHeader.Number.Uint64(),
		Timestamp:   blockHeader.Time,
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *WearablesFetcher) ParseWearableBurntEvent(rawLog etherTypes.Log) (wearables.WearableBurntEvent, error) {
	event, err := c.metronionAccessories.ParseAccessoryBurnt(rawLog)
	if err != nil {
		return wearables.WearableBurntEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableBurntEvent{}, err
	}

	return wearables.WearableBurntEvent{
		OnChainId:   event.Id.Uint64(),
		Amount:      event.Amount.Uint64(),
		From:        event.From.Hex(),
		BlockNumber: blockHeader.Number.Uint64(),
		Timestamp:   blockHeader.Time,
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *WearablesFetcher) ParseWearableStoredEvent(rawLog etherTypes.Log) (wearables.WearableStoredEvent, error) {
	event, err := c.metronionAccessories.ParseAccessoryStore(rawLog)
	if err != nil {
		return wearables.WearableStoredEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableStoredEvent{}, err
	}

	onChainIds := make([]uint64, len(event.AccessoryIds))
	for i := range event.AccessoryIds {
		onChainIds[i] = event.AccessoryIds[i].Uint64()
	}

	return wearables.WearableStoredEvent{
		OnChainIds:  onChainIds,
		From:        event.From.Hex(),
		BlockNumber: blockHeader.Number.Uint64(),
		Timestamp:   blockHeader.Time,
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *WearablesFetcher) ParseWearableReturnedEvent(rawLog etherTypes.Log) (wearables.WearbleReturnedEvent, error) {
	event, err := c.metronionAccessories.ParseAccessoryStore(rawLog)
	if err != nil {
		return wearables.WearbleReturnedEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearbleReturnedEvent{}, err
	}

	onChainIds := make([]uint64, len(event.AccessoryIds))
	for i := range event.AccessoryIds {
		onChainIds[i] = event.AccessoryIds[i].Uint64()
	}

	return wearables.WearbleReturnedEvent{
		OnChainIds:  onChainIds,
		From:        event.From.Hex(),
		BlockNumber: blockHeader.Number.Uint64(),
		Timestamp:   blockHeader.Time,
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *WearablesFetcher) ParseWearableTransferEvent(rawLog etherTypes.Log) (wearables.WearableTransferEvent, error) {
	event, err := c.metronionAccessories.ParseTransferSingle(rawLog)
	if err != nil {
		return wearables.WearableTransferEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableTransferEvent{}, err
	}

	return wearables.WearableTransferEvent{
		OnChainId:   event.Id.Uint64(),
		Amount:      event.Value.Uint64(),
		From:        event.From.Hex(),
		To:          event.To.Hex(),
		BlockNumber: blockHeader.Number.Uint64(),
		Timestamp:   blockHeader.Time,
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *WearablesFetcher) ParseWearableBatchTransferEvent(rawLog etherTypes.Log) (wearables.WearableBatchTransferEvent, error) {
	event, err := c.metronionAccessories.ParseTransferBatch(rawLog)
	if err != nil {
		return wearables.WearableBatchTransferEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableBatchTransferEvent{}, err
	}
	onChainIds := make([]uint64, len(event.Ids))
	for i := range event.Ids {
		onChainIds[i] = event.Ids[i].Uint64()
	}
	amounts := make([]uint64, len(event.Values))
	for i := range event.Values {
		amounts[i] = event.Values[i].Uint64()
	}

	return wearables.WearableBatchTransferEvent{
		OnChainIds:  onChainIds,
		Amounts:     amounts,
		From:        event.From.Hex(),
		To:          event.To.Hex(),
		BlockNumber: blockHeader.Number.Uint64(),
		Timestamp:   blockHeader.Time,
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func ParseWearablesType(input uint64) string {
	return services.MapWearablesType[input]
}

func ParseRarity(input uint8) string {
	return services.MapRarity[uint64(input)]
}
