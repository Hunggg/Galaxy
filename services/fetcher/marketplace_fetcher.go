package fetcher

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"

	etherCommon "github.com/ethereum/go-ethereum/common"
	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metrogalaxy-io/metrogalaxy-api/common"
	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/contracts"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/eventbus"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/wearables"
	"go.uber.org/zap"
)

type MarketplaceFetcher struct {
	l         *zap.SugaredLogger
	config    env.Config
	ethClient *ethclient.Client

	marketplaceAddress          etherCommon.Address
	metronionNFTAddress         etherCommon.Address
	metronionAccessoriesAddress etherCommon.Address
	marketplace                 *contracts.MetroGalaxyMarketplace
	topics                      []etherCommon.Hash
	listContracts               []etherCommon.Address
	cache                       cache.Cache
	eventFetcher                *EventFetcher
	subscriptionKeeper          *eventbus.SubscriptionKeeper
}

func NewMarketplaceFetcher(ethClient *ethclient.Client, eventFetcher *EventFetcher, cache cache.Cache, subscriptionKeeper *eventbus.SubscriptionKeeper) (*MarketplaceFetcher, error) {
	l := zap.S()
	config := env.GetConfig()

	marketplaceAddress := etherCommon.HexToAddress(config.MetroGalaxyMarketplaceContract.Address)
	marketplace, err := contracts.NewMetroGalaxyMarketplace(marketplaceAddress, ethClient)
	if err != nil {
		l.Errorw("initialize marketplace contract client error", "error", err)
		return nil, err
	}

	topics := []etherCommon.Hash{
		contracts.Marketplace_AssetListedTopic,
		contracts.Marketplace_AssetDelistedTopic,
		contracts.Marketplace_AssetOfferedTopic,
		contracts.Marketplace_AssetOfferCancelledTopic,
		contracts.Marketplace_AssetOfferTakenTopic,
		contracts.Marketplace_AssetBoughtTopic,
	}

	listContracts := []etherCommon.Address{
		marketplaceAddress,
	}

	return &MarketplaceFetcher{
		l:                           l,
		config:                      config,
		ethClient:                   ethClient,
		metronionNFTAddress:         etherCommon.HexToAddress(config.MetronionNFTContract.Address),
		metronionAccessoriesAddress: etherCommon.HexToAddress(config.MetronionAccessoriesContract.Address),
		marketplaceAddress:          marketplaceAddress,
		marketplace:                 marketplace,
		topics:                      topics,
		listContracts:               listContracts,
		cache:                       cache,
		eventFetcher:                eventFetcher,
		subscriptionKeeper:          subscriptionKeeper,
	}, nil
}

/**
 * Implementation of FetcherClient
 */
func (c *MarketplaceFetcher) FetchRawLogs(startBlock, endBlock, batchSize, numOfRoutines int64) ([]etherTypes.Log, error) {
	c.l.Debugw("marketplace fetcher fetch raw logs", "from_block", startBlock, "to_block", endBlock, "batch_size", batchSize)
	return c.eventFetcher.FetchBatchRawLogs(startBlock, endBlock, batchSize, c.topics, c.listContracts, numOfRoutines)
}

func (c *MarketplaceFetcher) FetcherClientName() FetcherClientName {
	return MetronionNFTFetcherClient
}

func (c *MarketplaceFetcher) Run(ctx context.Context, errCh chan<- error) {
	go c.SubscribeRawLogs(ctx, errCh)
}

/**
 * Parsing methods
 */
func (c *MarketplaceFetcher) SubscribeRawLogs(ctx context.Context, errCh chan<- error) {
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
			c.l.Errorw("subscribe metronion raw logs event error. Retry subscribing...", "error", err)
			sub = c.subscriptionKeeper.Subscribe(RawLogSubscriptionTopic, internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(etherTypes.Log)
			if !ok {
				c.l.Errorw("error decode raw logs event", "message", message)
				errCh <- fmt.Errorf("error decode raw logs event")
				return
			}
			if err := c.parseRawLogConcurrently(value); err != nil {
				c.l.Errorw("error parse raw logs event", "error", err)
				errCh <- fmt.Errorf("error parse raw logs event")
				return
			}
		}
	}
}

func (c *MarketplaceFetcher) parseRawLogConcurrently(rawLog etherTypes.Log) error {
	var (
		wg     = &sync.WaitGroup{}
		errCh  = make(chan error)
		doneCh = make(chan struct{})
	)

	go parseRawLogroutine(rawLog, wg, errCh, c.publishMetronionListingEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishMetronionDelistEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishMetronionBoughtEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishMetronionOfferEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishMetronionOfferCancelledEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishMetronionOfferTakenEvent)

	go parseRawLogroutine(rawLog, wg, errCh, c.publishWearableListingEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishWearableDelistEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishWearableOfferEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishWearableOfferCancelledEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishWearableOfferTakenEvent)
	go parseRawLogroutine(rawLog, wg, errCh, c.publishWearableBoughtEvent)

	go func() {
		wg.Wait()
		doneCh <- struct{}{}
		close(doneCh)
	}()

	for {
		select {
		case err := <-errCh:
			c.l.Errorw("error parse raw log", "error", err)
			return err
		case <-doneCh:
			return nil
		}
	}
}

func parseRawLogroutine(rawLog etherTypes.Log, wg *sync.WaitGroup, errCh chan<- error,
	handleLogFunc func(rawLog etherTypes.Log) error) {
	wg.Add(1)
	defer wg.Wait()
	err := handleLogFunc(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		errCh <- err
		return
	}
}

func (c *MarketplaceFetcher) publishMetronionListingEvent(rawLog etherTypes.Log) error {
	metronionListingEvent, err := c.ParseMetronionListing(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.MetronionListingEventId, metronionListingEvent)
	c.l.Infow("metronion listing event", "metronion_id", metronionListingEvent.MetronionID, "seller", metronionListingEvent.Seller, "price", common.WeiToFloat(metronionListingEvent.Price, 18), "timestamp", metronionListingEvent.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishMetronionDelistEvent(rawLog etherTypes.Log) error {
	metronionDelistEvent, err := c.ParseMetronionDelist(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.MetronionDelistEventId, metronionDelistEvent)
	c.l.Infow("metronion delist event", "metronion_id", metronionDelistEvent.MetronionID, "seller", metronionDelistEvent.Seller, "timestamp", metronionDelistEvent.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishMetronionOfferEvent(rawLog etherTypes.Log) error {
	metronionOfferEvent, err := c.ParseMetronionOffer(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.MetronionOfferEventId, metronionOfferEvent)
	c.l.Infow("metronion offer event", "metronion_id", metronionOfferEvent.MetronionID, "buyer", metronionOfferEvent.Buyer, "price", common.WeiToFloat(metronionOfferEvent.Price, 18), "timestamp", metronionOfferEvent.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishMetronionOfferCancelledEvent(rawLog etherTypes.Log) error {
	metronionOfferCancelledEvent, err := c.ParseMetronionOfferCancelled(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.MetronionOfferCancelledEventId, metronionOfferCancelledEvent)
	c.l.Infow("metronion offer cancelled event", "metronion_id", metronionOfferCancelledEvent.MetronionID, "buyer", metronionOfferCancelledEvent.Buyer, "timestamp", metronionOfferCancelledEvent.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishMetronionOfferTakenEvent(rawLog etherTypes.Log) error {
	metronionOfferTakenEvent, err := c.ParseMetronionOfferTaken(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.MetronionOfferTakenEventId, metronionOfferTakenEvent)
	c.l.Infow("metronion offer taken event", "metronion_id", metronionOfferTakenEvent.MetronionID, "buyer", metronionOfferTakenEvent.Buyer, "seller", metronionOfferTakenEvent.Seller, "price", common.WeiToFloat(metronionOfferTakenEvent.Price, 18), "timestamp", metronionOfferTakenEvent.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishMetronionBoughtEvent(rawLog etherTypes.Log) error {
	metronionBoughtEvent, err := c.ParseMetronionBought(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.MetronionBuyEventId, metronionBoughtEvent)
	c.l.Infow("metronion bought event", "metronion_id", metronionBoughtEvent.MetronionID, "buyer", metronionBoughtEvent.Buyer, "seller", metronionBoughtEvent.Seller, "price", common.WeiToFloat(metronionBoughtEvent.Price, 18), "timestamp", metronionBoughtEvent.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishWearableListingEvent(rawLog etherTypes.Log) error {
	event, err := c.ParseWearableListingEvent(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.WearableListingEventId, event)
	c.l.Infow("wearable listing event", "on_chain_id", event.OnChainId, "seller", event.Seller,
		"price", common.WeiToFloat(event.Price, 18), "amount", common.WeiToFloat(event.Amount, 0), "timestamp", event.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishWearableDelistEvent(rawLog etherTypes.Log) error {
	event, err := c.ParseWearableDelistEvent(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.WearableDelistEventId, event)
	c.l.Infow("wearable delist event", "on_chain_id", event.OnChainId, "seller", event.Seller,
		"timestamp", event.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishWearableOfferEvent(rawLog etherTypes.Log) error {
	event, err := c.ParseWearableOfferEvent(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.WearableOfferEventId, event)
	c.l.Infow("wearable offer event", "on_chain_id", event.OnChainId, "buyer", event.Buyer,
		"price", common.WeiToFloat(event.Price, 18), "amount", common.WeiToFloat(event.Amount, 0),
		"timestamp", event.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishWearableOfferCancelledEvent(rawLog etherTypes.Log) error {
	event, err := c.ParseWearableOfferCancelledEvent(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.WearableOfferCancelledEventId, event)
	c.l.Infow("wearable offer cancelled event", "on_chain_id", event.OnChainId, "buyer", event.Buyer,
		"timestamp", event.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishWearableOfferTakenEvent(rawLog etherTypes.Log) error {
	event, err := c.ParseWearableOfferTakenEvent(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.WearableOfferTakenEventId, event)
	c.l.Infow("wearable offer taken event", "on_chain_id", event.OnChainId, "buyer", event.Buyer,
		"seller", event.Seller, "price", common.WeiToFloat(event.Price, 18), "amount", common.WeiToFloat(event.Amount, 0),
		"timestamp", event.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) publishWearableBoughtEvent(rawLog etherTypes.Log) error {
	event, err := c.ParseWearableBoughtEvent(rawLog)
	if err != nil && err != common.ErrorCannotParseLog {
		return err
	}
	if err == common.ErrorCannotParseLog {
		return nil
	}
	c.subscriptionKeeper.Publish(contracts.WearableBoughtEventId, event)
	c.l.Infow("wearable bought event", "on_chain_id", event.OnChainId, "buyer", event.Buyer,
		"seller", event.Seller, "price", common.WeiToFloat(event.Price, 18), "amount", common.WeiToFloat(event.Amount, 0),
		"timestamp", event.Timestamp)
	return nil
}

func (c *MarketplaceFetcher) ParseMetronionListing(rawLog etherTypes.Log) (metronion.MetronionListingEvent, error) {
	event, err := c.marketplace.ParseAssetListed(rawLog)
	if err != nil {
		return metronion.MetronionListingEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionNFTAddress.Hex()) {
		return metronion.MetronionListingEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionListingEvent{}, err
	}

	return metronion.MetronionListingEvent{
		MetronionID: event.AssetId.Uint64(),
		Price:       event.PriceInWei,
		Seller:      strings.ToLower(event.Seller.Hex()),
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseMetronionDelist(rawLog etherTypes.Log) (metronion.MetronionDelistEvent, error) {
	event, err := c.marketplace.ParseAssetDelisted(rawLog)
	if err != nil {
		return metronion.MetronionDelistEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionNFTAddress.Hex()) {
		return metronion.MetronionDelistEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionDelistEvent{}, err
	}

	return metronion.MetronionDelistEvent{
		MetronionID: event.AssetId.Uint64(),
		Seller:      strings.ToLower(event.Seller.Hex()),
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseMetronionOffer(rawLog etherTypes.Log) (metronion.MetronionOfferEvent, error) {
	event, err := c.marketplace.ParseAssetOffered(rawLog)
	if err != nil {
		return metronion.MetronionOfferEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionNFTAddress.Hex()) {
		return metronion.MetronionOfferEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionOfferEvent{}, err
	}

	return metronion.MetronionOfferEvent{
		MetronionID: event.AssetId.Uint64(),
		Price:       event.PriceInWei,
		Buyer:       strings.ToLower(event.Buyer.Hex()),
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseMetronionOfferCancelled(rawLog etherTypes.Log) (metronion.MetronionOfferCancelledEvent, error) {
	event, err := c.marketplace.ParseAssetOfferCancelled(rawLog)
	if err != nil {
		return metronion.MetronionOfferCancelledEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionNFTAddress.Hex()) {
		return metronion.MetronionOfferCancelledEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionOfferCancelledEvent{}, err
	}

	return metronion.MetronionOfferCancelledEvent{
		MetronionID: event.AssetId.Uint64(),
		Buyer:       strings.ToLower(event.Buyer.Hex()),
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseMetronionOfferTaken(rawLog etherTypes.Log) (metronion.MetronionOfferTakenEvent, error) {
	event, err := c.marketplace.ParseAssetOfferTaken(rawLog)
	if err != nil {
		return metronion.MetronionOfferTakenEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionNFTAddress.Hex()) {
		return metronion.MetronionOfferTakenEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionOfferTakenEvent{}, err
	}

	return metronion.MetronionOfferTakenEvent{
		MetronionID: event.AssetId.Uint64(),
		Buyer:       strings.ToLower(event.Buyer.Hex()),
		Seller:      strings.ToLower(event.Seller.Hex()),
		Price:       event.PriceInWei,
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseMetronionBought(rawLog etherTypes.Log) (metronion.MetronionBuyEvent, error) {
	event, err := c.marketplace.ParseAssetBought(rawLog)
	if err != nil {
		return metronion.MetronionBuyEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionNFTAddress.Hex()) {
		return metronion.MetronionBuyEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionBuyEvent{}, err
	}

	return metronion.MetronionBuyEvent{
		MetronionID: event.AssetId.Uint64(),
		Buyer:       strings.ToLower(event.Buyer.Hex()),
		Seller:      strings.ToLower(event.Seller.Hex()),
		Price:       event.PriceInWei,
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseWearableListingEvent(rawLog etherTypes.Log) (wearables.WearableListingEvent, error) {
	event, err := c.marketplace.ParseAssetListed(rawLog)
	if err != nil {
		return wearables.WearableListingEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionAccessoriesAddress.Hex()) {
		return wearables.WearableListingEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableListingEvent{}, err
	}

	return wearables.WearableListingEvent{
		OnChainId:   event.AssetId.Uint64(),
		Price:       event.PriceInWei,
		Amount:      event.Amount,
		Seller:      strings.ToLower(event.Seller.Hex()),
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseWearableDelistEvent(rawLog etherTypes.Log) (wearables.WearableDelistEvent, error) {
	event, err := c.marketplace.ParseAssetDelisted(rawLog)
	if err != nil {
		return wearables.WearableDelistEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionAccessoriesAddress.Hex()) {
		return wearables.WearableDelistEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableDelistEvent{}, err
	}

	return wearables.WearableDelistEvent{
		OnChainId:   event.AssetId.Uint64(),
		Seller:      strings.ToLower(event.Seller.Hex()),
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseWearableOfferEvent(rawLog etherTypes.Log) (wearables.WearableOfferEvent, error) {
	event, err := c.marketplace.ParseAssetOffered(rawLog)
	if err != nil {
		return wearables.WearableOfferEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionAccessoriesAddress.Hex()) {
		return wearables.WearableOfferEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableOfferEvent{}, err
	}

	return wearables.WearableOfferEvent{
		OnChainId:   event.AssetId.Uint64(),
		Price:       event.PriceInWei,
		Amount:      event.Amount,
		Buyer:       strings.ToLower(event.Buyer.Hex()),
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseWearableOfferCancelledEvent(rawLog etherTypes.Log) (wearables.WearableOfferCancelledEvent, error) {
	event, err := c.marketplace.ParseAssetOfferCancelled(rawLog)
	if err != nil {
		return wearables.WearableOfferCancelledEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionAccessoriesAddress.Hex()) {
		return wearables.WearableOfferCancelledEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableOfferCancelledEvent{}, err
	}

	return wearables.WearableOfferCancelledEvent{
		OnChainId:   event.AssetId.Uint64(),
		Buyer:       strings.ToLower(event.Buyer.Hex()),
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseWearableOfferTakenEvent(rawLog etherTypes.Log) (wearables.WearableOfferTakenEvent, error) {
	event, err := c.marketplace.ParseAssetOfferTaken(rawLog)
	if err != nil {
		return wearables.WearableOfferTakenEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionAccessoriesAddress.Hex()) {
		return wearables.WearableOfferTakenEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableOfferTakenEvent{}, err
	}

	return wearables.WearableOfferTakenEvent{
		OnChainId:   event.AssetId.Uint64(),
		Buyer:       strings.ToLower(event.Buyer.Hex()),
		Seller:      strings.ToLower(event.Seller.Hex()),
		Price:       event.PriceInWei,
		Amount:      event.Amount,
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MarketplaceFetcher) ParseWearableBoughtEvent(rawLog etherTypes.Log) (wearables.WearableBoughtEvent, error) {
	event, err := c.marketplace.ParseAssetBought(rawLog)
	if err != nil {
		return wearables.WearableBoughtEvent{}, common.ErrorCannotParseLog
	}

	if !strings.EqualFold(event.AssetAddr.Hex(), c.metronionAccessoriesAddress.Hex()) {
		return wearables.WearableBoughtEvent{}, common.ErrorCannotParseLog
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return wearables.WearableBoughtEvent{}, err
	}

	return wearables.WearableBoughtEvent{
		OnChainId:   event.AssetId.Uint64(),
		Buyer:       strings.ToLower(event.Buyer.Hex()),
		Seller:      strings.ToLower(event.Seller.Hex()),
		Price:       event.PriceInWei,
		Amount:      event.Amount,
		Timestamp:   blockHeader.Time,
		BlockNumber: blockHeader.Number.Uint64(),
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}
