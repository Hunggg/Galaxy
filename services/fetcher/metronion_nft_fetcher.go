package fetcher

import (
	"context"
	"errors"
	"fmt"

	"math/big"
	"strings"

	"github.com/metrogalaxy-io/metrogalaxy-api/services"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	etherCommon "github.com/ethereum/go-ethereum/common"
	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/metrogalaxy-io/metrogalaxy-api/common"
	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/contracts"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/eventbus"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/cache"
	"go.uber.org/zap"
)

type MetronionNFTFetcher struct {
	l         *zap.SugaredLogger
	config    env.Config
	ethClient *ethclient.Client

	metronionNFTAddress etherCommon.Address
	metronionNFT        *contracts.MetronionNFT
	topics              []etherCommon.Hash
	listContracts       []etherCommon.Address
	metronionDb         metronion.PersistentDb
	cache               cache.Cache
	eventFetcher        *EventFetcher
	subscriptionKeeper  *eventbus.SubscriptionKeeper
}

func NewMetronionNFTFetcher(ethClient *ethclient.Client, eventFetcher *EventFetcher, metronionDb metronion.PersistentDb, cache cache.Cache, subscriptionKeeper *eventbus.SubscriptionKeeper) (*MetronionNFTFetcher, error) {
	l := zap.S()
	config := env.GetConfig()
	metronionNFT, err := contracts.NewMetronionNFT(etherCommon.HexToAddress(config.MetronionNFTContract.Address), ethClient)
	if err != nil {
		l.Errorw("initialize metronion nft contract client error", "error", err)
		return nil, err
	}

	topics := []etherCommon.Hash{
		contracts.MetronionNFT_TransferTopic,
		contracts.MetronionNFT_MetronionCreatedTopic,
		contracts.MetronionNFT_AccessoriesEquippedTopic,
		contracts.MetronionNFT_AccessoriesUnequippedTopic,
		contracts.MetronionNFT_NameChangedTopic,
		contracts.MetronionNFT_UpdateBaseURITopic,
	}

	listContracts := []etherCommon.Address{
		etherCommon.HexToAddress(config.MetronionNFTContract.Address),
	}

	return &MetronionNFTFetcher{
		l:                   l,
		config:              config,
		ethClient:           ethClient,
		metronionNFTAddress: etherCommon.HexToAddress(config.MetronionNFTContract.Address),
		metronionNFT:        metronionNFT,
		topics:              topics,
		listContracts:       listContracts,
		eventFetcher:        eventFetcher,
		metronionDb:         metronionDb,
		cache:               cache,
		subscriptionKeeper:  subscriptionKeeper,
	}, nil
}

/**
 * Implementation of FetcherClient
 */
func (c *MetronionNFTFetcher) FetchRawLogs(startBlock, endBlock, batchSize, numOfRoutines int64) ([]etherTypes.Log, error) {
	c.l.Debugw("metronion fetcher fetch raw logs", "from_block", startBlock, "to_block", endBlock, "batch_size", batchSize)
	return c.eventFetcher.FetchBatchRawLogs(startBlock, endBlock, batchSize, c.topics, c.listContracts, numOfRoutines)
}

func (c *MetronionNFTFetcher) FetcherClientName() FetcherClientName {
	return MetronionNFTFetcherClient
}

func (c *MetronionNFTFetcher) Run(ctx context.Context, errCh chan<- error) {
	// c.FetchVersionConfig()
	latestVersionId, err := c.FetchLatestVersionId()
	if err != nil {
		errCh <- err
		return
	}
	if err := c.FetchAndSaveVersionConfig(latestVersionId); err != nil {
		errCh <- err
		return
	}
	go c.SubscribeRawLogs(ctx, errCh)
}

/**
 * Internal Methods
 */
func (c *MetronionNFTFetcher) FetchLatestVersionId() (uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancel()
	resp, err := c.metronionNFT.GetLatestVersion(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		c.l.Errorw("fetch latest version id error", "error", err)
		return 0, err
	}
	return resp.Uint64(), nil
}

func (c *MetronionNFTFetcher) FetchAndSaveVersionConfig(versionID uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultContextTimeout)
	defer cancel()
	resp, err := c.metronionNFT.VersionById(&bind.CallOpts{
		Context: ctx,
	}, big.NewInt(int64(versionID)))
	if err != nil {
		c.l.Errorw("fetch version config error", "error", err, "version_id", versionID)
		return err
	}

	if err := c.metronionDb.SaveMetronionVersionConfig(metronion.MetronionVersionConfig{
		VersionID:     versionID,
		StartingIndex: resp.StartingIndex.Uint64(),
		MaxSupply:     resp.MaxSupply.Uint64(),
		CurrentSupply: resp.CurrentSupply.Uint64(),
		Provenance:    resp.Provenance,
	}); err != nil {
		c.l.Errorw("save version config error", "error", err, "version_id", versionID)
		return err
	}

	return nil
}

/**
 * Parsing methods
 */

func (c *MetronionNFTFetcher) SubscribeRawLogs(ctx context.Context, errCh chan<- error) {
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
func (c *MetronionNFTFetcher) ParseRawLog(rawLog etherTypes.Log) error {
	metronionCreatedEvent, err := c.ParseMetronionCreatedEvent(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionNFT_MetronionCreatedTopic.Hex(), metronionCreatedEvent)
		c.l.Infow("metronion created event", "metronion_id", metronionCreatedEvent.MetronionID, "owner", metronionCreatedEvent.Owner, "created_at_timestamp", metronionCreatedEvent.CreatedAtTimestamp)
		return nil
	}

	transferEvent, err := c.ParseMetronionTransferEvent(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionNFT_TransferTopic.Hex(), transferEvent)
		c.l.Infow("metronion transfer event", "metronion_id", transferEvent.MetronionID, "from", transferEvent.From, "to", transferEvent.To, "updated_at_timestamp", transferEvent.Timestamp)
		return nil
	}

	accessoriesEquippedEvent, err := c.ParseAccessoriesEquipped(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionNFT_AccessoriesEquippedTopic.Hex(), accessoriesEquippedEvent)
		c.l.Infow("metronion accessories equipped event", "metronion_id", accessoriesEquippedEvent.MetronionID, "accessory_ids", accessoriesEquippedEvent.AccessoryIds, "updated_at_timestamp", accessoriesEquippedEvent.Timestamp)
		return nil
	}

	accessoriesUnequippedEvent, err := c.ParseAccessoriesUnequipped(rawLog)
	if err == nil {
		c.subscriptionKeeper.Publish(contracts.MetronionNFT_AccessoriesUnequippedTopic.Hex(), accessoriesUnequippedEvent)
		c.l.Infow("metronion accessories unequipped event", "metronion_id", accessoriesUnequippedEvent.MetronionID, "accessory_ids", accessoriesUnequippedEvent.AccessoryIds, "updated_at_timestamp", accessoriesUnequippedEvent.Timestamp)
		return nil
	}
	return nil
}

func (c *MetronionNFTFetcher) ParseMetronionCreatedEvent(rawLog etherTypes.Log) (metronion.MetronionCreatedEvent, error) {
	metronionCreatedEvent, err := c.metronionNFT.ParseMetronionCreated(rawLog)
	if err != nil {
		return metronion.MetronionCreatedEvent{}, err
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(metronionCreatedEvent.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionCreatedEvent{}, err
	}

	return metronion.MetronionCreatedEvent{
		MetronionID:        metronionCreatedEvent.MetronionId.Uint64(),
		CreatedAtTimestamp: blockHeader.Time,
		UpdatedAtTimestamp: blockHeader.Time,
		CreatedAtBlock:     blockHeader.Number.Uint64(),
		UpdatedAtBlock:     blockHeader.Number.Uint64(),
		VersionID:          metronionCreatedEvent.VersionId.Uint64(),
		Owner:              strings.ToLower(metronionCreatedEvent.To.Hex()),
		BaseActivity: services.OnChainBaseActivity{
			From:        strings.ToLower(rawLog.Address.Hex()),
			To:          strings.ToLower(metronionCreatedEvent.To.Hex()),
			BlockNumber: blockHeader.Number.Uint64(),
			Timestamp:   blockHeader.Time,
			TxHash:      rawLog.TxHash.Hex(),
		},
	}, nil
}

func (c *MetronionNFTFetcher) ParseMetronionTransferEvent(rawLog etherTypes.Log) (metronion.MetronionTransferEvent, error) {
	transferEvent, err := c.metronionNFT.ParseTransfer(rawLog)
	if err != nil {
		return metronion.MetronionTransferEvent{}, err
	}
	if transferEvent.From == common.ZeroAddress {
		err = errors.New("invalid transfer event")
		return metronion.MetronionTransferEvent{}, err
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(transferEvent.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionTransferEvent{}, err
	}

	return metronion.MetronionTransferEvent{
		From:        strings.ToLower(transferEvent.From.Hex()),
		To:          strings.ToLower(transferEvent.To.Hex()),
		MetronionID: transferEvent.TokenId.Uint64(),
		BlockNumber: transferEvent.Raw.BlockNumber,
		Timestamp:   blockHeader.Time,
		TxHash:      rawLog.TxHash.Hex(),
	}, nil
}

func (c *MetronionNFTFetcher) ParseMetronionNameChangedEvent(rawLog etherTypes.Log) (metronion.MetronionNameChangedEvent, error) {
	nameChangedEvent, err := c.metronionNFT.ParseNameChanged(rawLog)
	if err != nil {
		return metronion.MetronionNameChangedEvent{}, err
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(nameChangedEvent.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionNameChangedEvent{}, err
	}

	return metronion.MetronionNameChangedEvent{
		MetronionID: nameChangedEvent.MetronionId.Uint64(),
		NewName:     nameChangedEvent.NewName,
		BlockNumber: nameChangedEvent.Raw.BlockNumber,
		Timestamp:   blockHeader.Time,
	}, nil
}

func (c *MetronionNFTFetcher) ParseAccessoriesEquipped(rawLog etherTypes.Log) (metronion.MetronionAccessoriesEquippedEvent, error) {
	event, err := c.metronionNFT.ParseAccessoriesEquipped(rawLog)
	if err != nil {
		return metronion.MetronionAccessoriesEquippedEvent{}, err
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionAccessoriesEquippedEvent{}, err
	}

	accessoryIds := make([]uint64, len(event.AccessoryIds))
	for _, item := range event.AccessoryIds {
		accessoryIds = append(accessoryIds, item.Uint64())
	}

	return metronion.MetronionAccessoriesEquippedEvent{
		MetronionID:  event.MetronionId.Uint64(),
		AccessoryIds: accessoryIds,
		BlockNumber:  event.Raw.BlockNumber,
		Timestamp:    blockHeader.Time,
	}, nil
}

func (c *MetronionNFTFetcher) ParseAccessoriesUnequipped(rawLog etherTypes.Log) (metronion.MetronionAccessoriesUnequippedEvent, error) {
	event, err := c.metronionNFT.ParseAccessoriesUnequipped(rawLog)
	if err != nil {
		return metronion.MetronionAccessoriesUnequippedEvent{}, err
	}

	blockHeader, err := common.GetBlockHeader(c.cache, c.ethClient, big.NewInt(int64(event.Raw.BlockNumber)))
	if err != nil {
		return metronion.MetronionAccessoriesUnequippedEvent{}, err
	}

	accessoryIds := make([]uint64, len(event.AccessoryIds))
	for _, item := range event.AccessoryIds {
		accessoryIds = append(accessoryIds, item.Uint64())
	}

	return metronion.MetronionAccessoriesUnequippedEvent{
		MetronionID:  event.MetronionId.Uint64(),
		AccessoryIds: accessoryIds,
		BlockNumber:  event.Raw.BlockNumber,
		Timestamp:    blockHeader.Time,
	}, nil
}
