package fetcher

import (
	"context"
	"fmt"

	"strings"

	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"

	"github.com/metrogalaxy-io/metrogalaxy-api/common"
	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/contracts"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/eventbus"
	"go.uber.org/zap"
)

type MetronionActivityHandler struct {
	l                  *zap.SugaredLogger
	config             env.Config
	metronionDb        metronion.PersistentDb
	subscriptionKeeper *eventbus.SubscriptionKeeper
}

func NewMetronionActivityHandler(metronionDb metronion.PersistentDb, subscriptionKeeper *eventbus.SubscriptionKeeper) *MetronionActivityHandler {
	return &MetronionActivityHandler{
		l:                  zap.S(),
		config:             env.GetConfig(),
		metronionDb:        metronionDb,
		subscriptionKeeper: subscriptionKeeper,
	}
}

func (h *MetronionActivityHandler) Run(ctx context.Context, errCh chan<- error) {
	go h.SubscribeMetronionCreatedEvent(ctx, errCh)
	go h.SubscribeMetronionTransfer(ctx, errCh)
	go h.SubscribeMetronionListing(ctx, errCh)
	go h.SubscribeMetronionDelist(ctx, errCh)
	go h.SubscribeMetronionOffer(ctx, errCh)
	go h.SubscribeMetronionOfferCancelled(ctx, errCh)
	go h.SubscribeMetronionOfferTaken(ctx, errCh)
	go h.SubscribeMetronionBuy(ctx, errCh)

	go h.SubscribeMetronionCurrentPrice(ctx, errCh)
	go h.SubscribeMetronionLastPrice(ctx, errCh)
}

func (h *MetronionActivityHandler) SubscribeMetronionCreatedEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_MetronionCreatedTopic.Hex(), internalErrCh)

	h.l.Infow("subscribe metronion created event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion created event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion created event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_MetronionCreatedTopic.Hex(), internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionCreatedEvent)
			if !ok {
				h.l.Errorw("error decode metronion created event", "message", message)
				errCh <- fmt.Errorf("error decode metronion created event")
				return
			}
			if err := h.metronionDb.SaveMetronionActivity(metronion.MetronionActivity{
				MetronionID:  value.MetronionID,
				ActivityType: metronion.ActivityTypeMint,
				FromAccount:  value.BaseActivity.From,
				ToAccount:    value.BaseActivity.To,
				BlockNumber:  value.BaseActivity.BlockNumber,
				Timestamp:    value.BaseActivity.Timestamp,
				TxHash:       value.BaseActivity.TxHash,
			}); err != nil {
				h.l.Errorw("error save metronion minted activity", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("save metronion minted activity", "metronion_id", value.MetronionID, "from", value.BaseActivity.From, "to", value.BaseActivity.To, "timestamp", value.BaseActivity.Timestamp)
		}
	}
}

func (h *MetronionActivityHandler) SubscribeMetronionTransfer(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_TransferTopic.Hex(), internalErrCh)

	h.l.Infow("subscribe metronion transfer event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion transfer event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion transfer event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_TransferTopic.Hex(), internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionTransferEvent)
			if !ok {
				h.l.Errorw("error decode metronion transfer event", "message", message)
				errCh <- fmt.Errorf("error decode metronion transfer event")
				return
			}
			if strings.EqualFold(value.From, h.config.MetroGalaxyMarketplaceContract.Address) || strings.EqualFold(value.To, h.config.MetroGalaxyMarketplaceContract.Address) {
				continue
			}
			if err := h.metronionDb.SaveMetronionActivity(metronion.MetronionActivity{
				MetronionID:  value.MetronionID,
				ActivityType: metronion.ActivityTypeTransfer,
				FromAccount:  value.From,
				ToAccount:    value.To,
				BlockNumber:  value.BlockNumber,
				Timestamp:    value.Timestamp,
				TxHash:       value.TxHash,
			}); err != nil {
				h.l.Errorw("error update metronion transfer activity", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion transfer activity", "metronion_id", value.MetronionID, "from", value.From, "to", value.To, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionActivityHandler) SubscribeMetronionListing(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionListingEventId, internalErrCh)

	h.l.Infow("subscribe metronion listing event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion listing event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion listing event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionListingEventId, internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionListingEvent)
			if !ok {
				h.l.Errorw("error decode metronion listing event", "message", message)
				errCh <- fmt.Errorf("error decode metronion listing event")
				return
			}

			price := common.WeiToFloat(value.Price, 18)
			if err := h.metronionDb.SaveMetronionActivity(metronion.MetronionActivity{
				MetronionID:  value.MetronionID,
				ActivityType: metronion.ActivityTypeList,
				FromAccount:  value.Seller,
				Price:        price,
				BlockNumber:  value.BlockNumber,
				Timestamp:    value.Timestamp,
				TxHash:       value.TxHash,
			}); err != nil {
				h.l.Errorw("error update metronion listing activity", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion listing activity", "metronion_id", value.MetronionID, "seller", value.Seller, "price", price, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionActivityHandler) SubscribeMetronionDelist(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionDelistEventId, internalErrCh)

	h.l.Infow("subscribe metronion delist event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion delist event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion delist event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionDelistEventId, internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionDelistEvent)
			if !ok {
				h.l.Errorw("error decode metronion delist event", "message", message)
				errCh <- fmt.Errorf("error decode metronion delist event")
				return
			}
			if err := h.metronionDb.SaveMetronionActivity(metronion.MetronionActivity{
				MetronionID:  value.MetronionID,
				ActivityType: metronion.ActivityTypeDelist,
				FromAccount:  value.Seller,
				BlockNumber:  value.BlockNumber,
				Timestamp:    value.Timestamp,
				TxHash:       value.TxHash,
			}); err != nil {
				h.l.Errorw("error update metronion delist activity", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion delist activity", "metronion_id", value.MetronionID, "seller", value.Seller, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionActivityHandler) SubscribeMetronionOffer(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionOfferEventId, internalErrCh)

	h.l.Infow("subscribe metronion offer event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion offer event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion offer event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionOfferEventId, internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionOfferEvent)
			if !ok {
				h.l.Errorw("error decode metronion offer event", "message", message)
				errCh <- fmt.Errorf("error decode metronion offer event")
				return
			}
			price := common.WeiToFloat(value.Price, 18)
			if err := h.metronionDb.SaveMetronionActivity(metronion.MetronionActivity{
				MetronionID:  value.MetronionID,
				ActivityType: metronion.ActivityTypeOffer,
				FromAccount:  value.Buyer,
				Price:        price,
				BlockNumber:  value.BlockNumber,
				Timestamp:    value.Timestamp,
				TxHash:       value.TxHash,
			}); err != nil {
				h.l.Errorw("error update metronion offer activity", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion offer activity", "metronion_id", value.MetronionID, "buyer", value.Buyer, "price", price, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionActivityHandler) SubscribeMetronionOfferCancelled(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionOfferCancelledEventId, internalErrCh)

	h.l.Infow("subscribe metronion offer cancelled event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion offer cancelled event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion offer cancelled event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionOfferCancelledEventId, internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionOfferCancelledEvent)
			if !ok {
				h.l.Errorw("error decode metronion offer cancelled event", "message", message)
				errCh <- fmt.Errorf("error decode metronion offer cancelled event")
				return
			}
			if err := h.metronionDb.SaveMetronionActivity(metronion.MetronionActivity{
				MetronionID:  value.MetronionID,
				ActivityType: metronion.ActivityTypeOfferCancelled,
				FromAccount:  value.Buyer,
				BlockNumber:  value.BlockNumber,
				Timestamp:    value.Timestamp,
				TxHash:       value.TxHash,
			}); err != nil {
				h.l.Errorw("error update metronion offer cancelled activity", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion offer cancelled activity", "metronion_id", value.MetronionID, "buyer", value.Buyer, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionActivityHandler) SubscribeMetronionOfferTaken(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionOfferTakenEventId, internalErrCh)

	h.l.Infow("subscribe metronion offer taken event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion offer taken event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion offer taken event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionOfferTakenEventId, internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionOfferTakenEvent)
			if !ok {
				h.l.Errorw("error decode metronion offer taken event", "message", message)
				errCh <- fmt.Errorf("error decode metronion offer taken event")
				return
			}
			price := common.WeiToFloat(value.Price, 18)
			if err := h.metronionDb.SaveMetronionActivity(metronion.MetronionActivity{
				MetronionID:  value.MetronionID,
				ActivityType: metronion.ActivityTypeOfferTaken,
				FromAccount:  value.Seller,
				ToAccount:    value.Buyer,
				Price:        price,
				BlockNumber:  value.BlockNumber,
				Timestamp:    value.Timestamp,
				TxHash:       value.TxHash,
			}); err != nil {
				h.l.Errorw("error update metronion offer taken activity", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion offer taken activity", "metronion_id", value.MetronionID, "seller", value.Seller, "buyer", value.Buyer, "price", price, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionActivityHandler) SubscribeMetronionBuy(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionBuyEventId, internalErrCh)

	h.l.Infow("subscribe metronion buy event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion buy event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion offer buy error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionBuyEventId, internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionBuyEvent)
			if !ok {
				h.l.Errorw("error decode metronion buy event", "message", message)
				errCh <- fmt.Errorf("error decode metronion buy event")
				return
			}
			price := common.WeiToFloat(value.Price, 18)
			if err := h.metronionDb.SaveMetronionActivity(metronion.MetronionActivity{
				MetronionID:  value.MetronionID,
				ActivityType: metronion.ActivityTypeSale,
				FromAccount:  value.Seller,
				ToAccount:    value.Buyer,
				Price:        price,
				BlockNumber:  value.BlockNumber,
				Timestamp:    value.Timestamp,
				TxHash:       value.TxHash,
			}); err != nil {
				h.l.Errorw("error update metronion buy activity", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion buy activity", "metronion_id", value.MetronionID, "seller", value.Seller, "buyer", value.Buyer, "price", price, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionActivityHandler) SubscribeMetronionCurrentPrice(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)

	metronionListingSub := h.subscriptionKeeper.Subscribe(contracts.MetronionListingEventId, internalErrCh)
	metronionDelistSub := h.subscriptionKeeper.Subscribe(contracts.MetronionDelistEventId, internalErrCh)
	metronionBoughtSub := h.subscriptionKeeper.Subscribe(contracts.MetronionBuyEventId, internalErrCh)
	metronionOfferTakenSub := h.subscriptionKeeper.Subscribe(contracts.MetronionOfferTakenEventId, internalErrCh)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion listing event", "sub_id", metronionListingSub.Id)
			h.subscriptionKeeper.Unsubscribe(metronionListingSub.Id)
			h.l.Debugw("unsubscribe metronion delist event", "sub_id", metronionDelistSub.Id)
			h.subscriptionKeeper.Unsubscribe(metronionDelistSub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion current price error. Retry subscribing...", "error", err)
			metronionListingSub = h.subscriptionKeeper.Subscribe(contracts.MetronionListingEventId, internalErrCh)
			metronionDelistSub = h.subscriptionKeeper.Subscribe(contracts.MetronionDelistEventId, internalErrCh)
		case message := <-metronionListingSub.Receiver:
			value, ok := message.(metronion.MetronionListingEvent)
			if !ok {
				h.l.Errorw("error decode metronion listing event", "message", message)
				errCh <- fmt.Errorf("error decode metronion listing event")
				return
			}

			price := common.WeiToFloat(value.Price, 18)
			if err := h.metronionDb.UpdateMetronionPrice(value.MetronionID, metronion.MetronionPriceTypeCurrentPrice, price); err != nil {
				h.l.Errorw("update metronion current price error", "error", err)
				return
			}
		case message := <-metronionDelistSub.Receiver:
			value, ok := message.(metronion.MetronionDelistEvent)
			if !ok {
				h.l.Errorw("error decode metronion delist event", "message", message)
				errCh <- fmt.Errorf("error decode metronion delist event")
				return
			}

			if err := h.metronionDb.UpdateMetronionPrice(value.MetronionID, metronion.MetronionPriceTypeCurrentPrice, 0); err != nil {
				h.l.Errorw("update metronion current price error", "error", err)
				return
			}
		case message := <-metronionBoughtSub.Receiver:
			value, ok := message.(metronion.MetronionBuyEvent)
			if !ok {
				h.l.Errorw("error decode metronion buy event", "message", message)
				errCh <- fmt.Errorf("error decode metronion buy event")
				return
			}

			if err := h.metronionDb.UpdateMetronionPrice(value.MetronionID, metronion.MetronionPriceTypeCurrentPrice, 0); err != nil {
				h.l.Errorw("update metronion current price error", "error", err)
				return
			}
		case message := <-metronionOfferTakenSub.Receiver:
			value, ok := message.(metronion.MetronionOfferTakenEvent)
			if !ok {
				h.l.Errorw("error decode metronion offer taken event", "message", message)
				errCh <- fmt.Errorf("error decode metronion offer taken event")
				return
			}
			if err := h.metronionDb.UpdateMetronionPrice(value.MetronionID, metronion.MetronionPriceTypeCurrentPrice, 0); err != nil {
				h.l.Errorw("update metronion current price error", "error", err)
				return
			}
		}
	}
}

func (h *MetronionActivityHandler) SubscribeMetronionLastPrice(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)

	metronionOfferTakenSub := h.subscriptionKeeper.Subscribe(contracts.MetronionOfferTakenEventId, internalErrCh)
	metronionBoughtSub := h.subscriptionKeeper.Subscribe(contracts.MetronionBuyEventId, internalErrCh)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion offer taken event", "sub_id", metronionOfferTakenSub.Id)
			h.subscriptionKeeper.Unsubscribe(metronionOfferTakenSub.Id)
			h.l.Debugw("unsubscribe metronion bought event", "sub_id", metronionBoughtSub.Id)
			h.subscriptionKeeper.Unsubscribe(metronionBoughtSub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion last price error. Retry subscribing...", "error", err)
			metronionOfferTakenSub = h.subscriptionKeeper.Subscribe(contracts.MetronionOfferTakenEventId, internalErrCh)
			metronionBoughtSub = h.subscriptionKeeper.Subscribe(contracts.MetronionBuyEventId, internalErrCh)
		case message := <-metronionOfferTakenSub.Receiver:
			value, ok := message.(metronion.MetronionOfferTakenEvent)
			if !ok {
				h.l.Errorw("error decode metronion offer taken event", "message", message)
				errCh <- fmt.Errorf("error decode metronion offer taken event")
				return
			}
			price := common.WeiToFloat(value.Price, 18)
			if err := h.metronionDb.UpdateMetronionPrice(value.MetronionID, metronion.MetronionPriceTypeLastPrice, price); err != nil {
				h.l.Errorw("update metronion last price error", "error", err)
				return
			}
		case message := <-metronionBoughtSub.Receiver:
			value, ok := message.(metronion.MetronionBuyEvent)
			if !ok {
				h.l.Errorw("error decode metronion bought event", "message", message)
				errCh <- fmt.Errorf("error decode metronion bought event")
				return
			}
			price := common.WeiToFloat(value.Price, 18)
			if err := h.metronionDb.UpdateMetronionPrice(value.MetronionID, metronion.MetronionPriceTypeLastPrice, price); err != nil {
				h.l.Errorw("update metronion last price error", "error", err)
				return
			}
		}
	}
}
