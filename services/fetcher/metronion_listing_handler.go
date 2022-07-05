package fetcher

import (
	"context"
	"fmt"

	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"

	"github.com/metrogalaxy-io/metrogalaxy-api/common"
	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/contracts"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/eventbus"
	"go.uber.org/zap"
)

type MetronionListingHandler struct {
	l                  *zap.SugaredLogger
	config             env.Config
	metronionDb        metronion.PersistentDb
	subscriptionKeeper *eventbus.SubscriptionKeeper
}

func NewMetronionListingHandler(metronionDb metronion.PersistentDb, subscriptionKeeper *eventbus.SubscriptionKeeper) *MetronionListingHandler {
	return &MetronionListingHandler{
		l:                  zap.S(),
		config:             env.GetConfig(),
		metronionDb:        metronionDb,
		subscriptionKeeper: subscriptionKeeper,
	}
}

func (h *MetronionListingHandler) Run(ctx context.Context, errCh chan<- error) {
	go h.SubscribeMetronionListing(ctx, errCh)
	go h.SubscribeMetronionDelist(ctx, errCh)
	go h.SubscribeMetronionSale(ctx, errCh)
}

func (h *MetronionListingHandler) SubscribeMetronionListing(ctx context.Context, errCh chan<- error) {
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
			if err := h.metronionDb.UpdateMetronionListing(metronion.MetronionListing{
				MetronionID: value.MetronionID,
				FromAccount: value.Seller,
				Price:       price,
				BlockNumber: value.BlockNumber,
				Timestamp:   value.Timestamp,
				TxHash:      value.TxHash,
			}, true); err != nil {
				h.l.Errorw("error update metronion listing", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("add metronion listing", "metronion_id", value.MetronionID, "seller", value.Seller, "price", price, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionListingHandler) SubscribeMetronionDelist(ctx context.Context, errCh chan<- error) {
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
			if err := h.metronionDb.UpdateMetronionListing(metronion.MetronionListing{
				MetronionID: value.MetronionID,
				FromAccount: value.Seller,
				BlockNumber: value.BlockNumber,
				Timestamp:   value.Timestamp,
				TxHash:      value.TxHash,
			}, false); err != nil {
				h.l.Errorw("error update metronion delist", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("remove metronion listing", "metronion_id", value.MetronionID, "seller", value.Seller, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionListingHandler) SubscribeMetronionSale(ctx context.Context, errCh chan<- error) {
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
			if err := h.metronionDb.UpdateMetronionListing(metronion.MetronionListing{
				MetronionID: value.MetronionID,
				FromAccount: value.Seller,
				BlockNumber: value.BlockNumber,
				Timestamp:   value.Timestamp,
				TxHash:      value.TxHash,
			}, false); err != nil {
				h.l.Errorw("update metronion delist error", "error", err)
				return
			}
			h.l.Debugw("remove metronion listing", "metronion_id", value.MetronionID, "seller", value.Seller, "timestamp", value.Timestamp)
		case message := <-metronionBoughtSub.Receiver:
			value, ok := message.(metronion.MetronionBuyEvent)
			if !ok {
				h.l.Errorw("error decode metronion bought event", "message", message)
				errCh <- fmt.Errorf("error decode metronion bought event")
				return
			}
			if err := h.metronionDb.UpdateMetronionListing(metronion.MetronionListing{
				MetronionID: value.MetronionID,
				FromAccount: value.Seller,
				BlockNumber: value.BlockNumber,
				Timestamp:   value.Timestamp,
				TxHash:      value.TxHash,
			}, false); err != nil {
				h.l.Errorw("update metronion delist error", "error", err)
				return
			}
			h.l.Debugw("remove metronion listing", "metronion_id", value.MetronionID, "seller", value.Seller, "timestamp", value.Timestamp)
		}
	}
}
