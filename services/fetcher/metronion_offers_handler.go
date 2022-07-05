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

type MetronionOffersHandler struct {
	l                  *zap.SugaredLogger
	config             env.Config
	metronionDb        metronion.PersistentDb
	subscriptionKeeper *eventbus.SubscriptionKeeper
}

func NewMetronionOffersHandler(metronionDb metronion.PersistentDb, subscriptionKeeper *eventbus.SubscriptionKeeper) *MetronionOffersHandler {
	return &MetronionOffersHandler{
		l:                  zap.S(),
		config:             env.GetConfig(),
		metronionDb:        metronionDb,
		subscriptionKeeper: subscriptionKeeper,
	}
}

func (h *MetronionOffersHandler) Run(ctx context.Context, errCh chan<- error) {
	go h.SubscribeMetronionOffer(ctx, errCh)
	go h.SubscribeMetronionOfferCancelled(ctx, errCh)
	go h.SubscribeMetronionOfferTaken(ctx, errCh)
}

func (h *MetronionOffersHandler) SubscribeMetronionOffer(ctx context.Context, errCh chan<- error) {
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
			if err := h.metronionDb.UpdateMetronionOffers(metronion.MetronionOffers{
				MetronionID: value.MetronionID,
				FromAccount: value.Buyer,
				Price:       price,
				BlockNumber: value.BlockNumber,
				Timestamp:   value.Timestamp,
				TxHash:      value.TxHash,
			}, true); err != nil {
				h.l.Errorw("error update metronion offer", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("add metronion offers", "metronion_id", value.MetronionID, "buyer", value.Buyer, "price", price, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionOffersHandler) SubscribeMetronionOfferCancelled(ctx context.Context, errCh chan<- error) {
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
			if err := h.metronionDb.UpdateMetronionOffers(metronion.MetronionOffers{
				MetronionID: value.MetronionID,
				FromAccount: value.Buyer,
				BlockNumber: value.BlockNumber,
				Timestamp:   value.Timestamp,
				TxHash:      value.TxHash,
			}, false); err != nil {
				h.l.Errorw("error update metronion offer", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("remove metronion offers", "metronion_id", value.MetronionID, "buyer", value.Buyer, "timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionOffersHandler) SubscribeMetronionOfferTaken(ctx context.Context, errCh chan<- error) {
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
			if err := h.metronionDb.UpdateMetronionOffers(metronion.MetronionOffers{
				MetronionID: value.MetronionID,
				FromAccount: value.Buyer,
				Price:       price,
				BlockNumber: value.BlockNumber,
				Timestamp:   value.Timestamp,
				TxHash:      value.TxHash,
			}, false); err != nil {
				h.l.Errorw("error update metronion offer", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("remove metronion offers", "metronion_id", value.MetronionID, "seller", value.Seller, "buyer", value.Buyer, "price", price, "timestamp", value.Timestamp)
		}
	}
}
