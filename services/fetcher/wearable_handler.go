package fetcher

import (
	"context"
	"fmt"

	"github.com/metrogalaxy-io/metrogalaxy-api/common"
	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/contracts"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/eventbus"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/wearables"
	"go.uber.org/zap"
)

type WearableHandler struct {
	l                  *zap.SugaredLogger
	config             env.Config
	db                 wearables.PersistentDb
	subscriptionKeeper *eventbus.SubscriptionKeeper
}

func NewWearableHandler(db wearables.PersistentDb, subscriptionKeeper *eventbus.SubscriptionKeeper) *WearableHandler {
	return &WearableHandler{
		l:                  zap.S(),
		config:             env.GetConfig(),
		db:                 db,
		subscriptionKeeper: subscriptionKeeper,
	}
}

func (h *WearableHandler) Run(ctx context.Context, errCh chan<- error) {
	go h.subscribeWearableCreatedEvent(ctx, errCh)
	go h.subscribeWearableMintedEvent(ctx, errCh)
	go h.subscribeWearableBurntEvent(ctx, errCh)
	go h.subscribeWearableListingEvent(ctx, errCh)
	go h.subscribeWearableDelistEvent(ctx, errCh)
	go h.subscribeWearableOfferEvent(ctx, errCh)
	go h.subscribeWearableOfferCancelledEvent(ctx, errCh)
	go h.subscribeWearableOfferTakenEvent(ctx, errCh)
	go h.subscribeWearableBoughtEvent(ctx, errCh)
}

func (h *WearableHandler) subscribeWearableCreatedEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionAccessories_AccessoryCreatedTopic.Hex(), internalErrCh)

	h.l.Infow("subscribe wearable created event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe wearable created event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe wearable created event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionAccessories_AccessoryCreatedTopic.Hex(), internalErrCh)
		case message := <-sub.Receiver:
			event, ok := message.(wearables.WearableCreatedEvent)
			if !ok {
				h.l.Errorw("error decode wearable created event", "message", message)
				errCh <- fmt.Errorf("error decode wearable created event")
				return
			}
			if err := h.saveWearableCreated(event); err != nil {
				h.l.Errorw("error save wearable", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("save wearable created", "wearable_on_chain_id", event.OnChainId, "name", event.Name,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		}
	}
}

func (h *WearableHandler) subscribeWearableMintedEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionAccessories_AccessoryMintTopic.Hex(), internalErrCh)

	h.l.Infow("subscribe wearable minted event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe wearable minted event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe wearable minted event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionAccessories_AccessoryMintTopic.Hex(), internalErrCh)
		case message := <-sub.Receiver:
			event, ok := message.(wearables.WearableMintedEvent)
			if !ok {
				h.l.Errorw("error decode wearable minted event", "message", message)
				errCh <- fmt.Errorf("error decode wearable minted event")
				return
			}
			if err := h.updateWearableMinted(event); err != nil {
				h.l.Errorw("error save wearable", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update wearable minted", "wearable_on_chain_id", event.OnChainId,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		}
	}
}

func (h *WearableHandler) subscribeWearableBurntEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionAccessories_AccessoryBurntTopic.Hex(), internalErrCh)

	h.l.Infow("subscribe wearable burnt event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe wearable burnt event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe wearable burnt event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionAccessories_AccessoryMintTopic.Hex(), internalErrCh)
		case message := <-sub.Receiver:
			event, ok := message.(wearables.WearableBurntEvent)
			if !ok {
				h.l.Errorw("error decode wearable burnt event", "message", message)
				errCh <- fmt.Errorf("error decode wearable burnt event")
				return
			}
			if err := h.updateWearableBurnt(event); err != nil {
				h.l.Errorw("error update wearable", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update wearable burnt", "wearable_on_chain_id", event.OnChainId,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		}
	}
}

func (h *WearableHandler) subscribeWearableListingEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.WearableListingEventId, internalErrCh)

	h.l.Infow("subscribe wearable listing event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe wearable listing event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe wearable listing event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.WearableListingEventId, internalErrCh)
		case message := <-sub.Receiver:
			event, ok := message.(wearables.WearableListingEvent)
			if !ok {
				h.l.Errorw("error decode wearable listing event", "message", message)
				errCh <- fmt.Errorf("error decode wearable listing event")
				return
			}
			price := common.WeiToFloat(event.Price, 18)
			amount := common.WeiToFloat(event.Amount, 0)
			if err := h.db.UpdateWearableListing(wearables.WearableListing{
				OnChainId:   event.OnChainId,
				FromAccount: event.Seller,
				Price:       price,
				Amount:      uint64(amount),
				Timestamp:   event.Timestamp,
				BlockNumber: event.BlockNumber,
				TxHash:      event.TxHash,
			}, true); err != nil {
				h.l.Errorw("error update wearable listing", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("add wearable listing", "wearable_on_chain_id", event.OnChainId,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		}
	}
}

func (h *WearableHandler) subscribeWearableDelistEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.WearableDelistEventId, internalErrCh)

	h.l.Infow("subscribe wearable delist event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe wearable delist event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe wearable delist event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.WearableDelistEventId, internalErrCh)
		case message := <-sub.Receiver:
			event, ok := message.(wearables.WearableDelistEvent)
			if !ok {
				h.l.Errorw("error decode wearable delist event", "message", message)
				errCh <- fmt.Errorf("error decode wearable delist event")
				return
			}
			if err := h.db.UpdateWearableListing(wearables.WearableListing{
				OnChainId:   event.OnChainId,
				FromAccount: event.Seller,
				Timestamp:   event.Timestamp,
				BlockNumber: event.BlockNumber,
				TxHash:      event.TxHash}, false); err != nil {
				h.l.Errorw("error update wearable delist", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("remove wearable listing", "wearable_on_chain_id", event.OnChainId,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		}
	}
}

func (h *WearableHandler) subscribeWearableOfferEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.WearableOfferEventId, internalErrCh)

	h.l.Infow("subscribe wearable offer event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe wearable offer event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe wearable offer event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.WearableOfferEventId, internalErrCh)
		case message := <-sub.Receiver:
			event, ok := message.(wearables.WearableOfferEvent)
			if !ok {
				h.l.Errorw("error decode wearable offer event", "message", message)
				errCh <- fmt.Errorf("error decode wearable offer event")
				return
			}
			price := common.WeiToFloat(event.Price, 18)
			amount := common.WeiToFloat(event.Amount, 0)
			if err := h.db.UpdateWearableOffer(wearables.WearableOffers{
				OnChainId:   event.OnChainId,
				FromAccount: event.Buyer,
				Price:       price,
				Amount:      uint64(amount),
				Timestamp:   event.Timestamp,
				BlockNumber: event.BlockNumber,
				TxHash:      event.TxHash,
			}, true); err != nil {
				h.l.Errorw("error update wearable offer", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("add wearable offer", "wearable_on_chain_id", event.OnChainId,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		}
	}
}

func (h *WearableHandler) subscribeWearableOfferCancelledEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.WearableOfferCancelledEventId, internalErrCh)

	h.l.Infow("subscribe wearable offer cancelled event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe wearable offer cancelled event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe wearable offer cancelled event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.WearableOfferCancelledEventId, internalErrCh)
		case message := <-sub.Receiver:
			event, ok := message.(wearables.WearableOfferCancelledEvent)
			if !ok {
				h.l.Errorw("error decode wearable offer cancelled event", "message", message)
				errCh <- fmt.Errorf("error decode wearable offer cancelled event")
				return
			}
			if err := h.db.UpdateWearableOffer(wearables.WearableOffers{
				OnChainId:   event.OnChainId,
				FromAccount: event.Buyer,
				Timestamp:   event.Timestamp,
				BlockNumber: event.BlockNumber,
				TxHash:      event.TxHash,
			}, false); err != nil {
				h.l.Errorw("error update wearable offer cancelled", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("remove wearable offer", "wearable_on_chain_id", event.OnChainId,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		}
	}
}

func (h *WearableHandler) subscribeWearableOfferTakenEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.WearableOfferTakenEventId, internalErrCh)

	h.l.Infow("subscribe wearable offer taken event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe wearable offer taken event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe wearable offer taken event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.WearableOfferTakenEventId, internalErrCh)
		case message := <-sub.Receiver:
			event, ok := message.(wearables.WearableOfferTakenEvent)
			if !ok {
				h.l.Errorw("error decode wearable offer taken event", "message", message)
				errCh <- fmt.Errorf("error decode wearable offer taken event")
				return
			}
			if err := h.db.UpdateWearableOffer(wearables.WearableOffers{
				OnChainId:   event.OnChainId,
				FromAccount: event.Buyer,
				Timestamp:   event.Timestamp,
				BlockNumber: event.BlockNumber,
				TxHash:      event.TxHash,
			}, false); err != nil {
				h.l.Errorw("error update wearable offer taken", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("remove wearable offer", "wearable_on_chain_id", event.OnChainId,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		}
	}
}

func (h *WearableHandler) subscribeWearableBoughtEvent(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)

	wearableBoughtSub := h.subscriptionKeeper.Subscribe(contracts.WearableBoughtEventId, internalErrCh)
	wearableOfferTakenSub := h.subscriptionKeeper.Subscribe(contracts.WearableOfferTakenEventId, internalErrCh)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe wearable offer taken event", "sub_id", wearableOfferTakenSub.Id)
			h.subscriptionKeeper.Unsubscribe(wearableOfferTakenSub.Id)
			h.l.Debugw("unsubscribe wearable bought event", "sub_id", wearableBoughtSub.Id)
			h.subscriptionKeeper.Unsubscribe(wearableBoughtSub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion last price error. Retry subscribing...", "error", err)
			wearableOfferTakenSub = h.subscriptionKeeper.Subscribe(contracts.WearableOfferTakenEventId, internalErrCh)
			wearableBoughtSub = h.subscriptionKeeper.Subscribe(contracts.WearableBoughtEventId, internalErrCh)
		case message := <-wearableOfferTakenSub.Receiver:
			event, ok := message.(wearables.WearableOfferTakenEvent)
			if !ok {
				h.l.Errorw("error decode wearable offer taken event", "message", message)
				errCh <- fmt.Errorf("error decode wearable offer taken event")
				return
			}
			if err := h.db.UpdateWearableListing(wearables.WearableListing{
				OnChainId:   event.OnChainId,
				FromAccount: event.Buyer,
				Timestamp:   event.Timestamp,
				BlockNumber: event.BlockNumber,
				TxHash:      event.TxHash,
			}, false); err != nil {
				h.l.Errorw("error remove wearable listing", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("remove wearable listing", "wearable_on_chain_id", event.OnChainId,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		case message := <-wearableBoughtSub.Receiver:
			event, ok := message.(wearables.WearableBoughtEvent)
			if !ok {
				h.l.Errorw("error decode wearable bought event", "message", message)
				errCh <- fmt.Errorf("error decode wearable bought event")
				return
			}
			if err := h.db.UpdateWearableListing(wearables.WearableListing{
				OnChainId:   event.OnChainId,
				FromAccount: event.Buyer,
				Timestamp:   event.Timestamp,
				BlockNumber: event.BlockNumber,
				TxHash:      event.TxHash,
			}, false); err != nil {
				h.l.Errorw("error remove wearable listing", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("remove wearable listing", "wearable_on_chain_id", event.OnChainId,
				"timestamp", event.Timestamp, "block_number", event.BlockNumber)
		}
	}
}

func (h *WearableHandler) saveWearableCreated(event wearables.WearableCreatedEvent) error {
	data := wearables.WearablesOnChainData{
		Id:                 event.OnChainId,
		Name:               event.Name,
		MaxSupply:          event.MaxSupply,
		CurrentSupply:      0,
		Type:               event.Type,
		Rarity:             event.Rarity,
		CreatedAtTimestamp: event.Timestamp,
		UpdatedAtTimestamp: event.Timestamp,
		CreatedAtBlock:     event.BlockNumber,
		UpdatedAtBlock:     event.BlockNumber,
	}
	if err := h.db.SaveWearableOnChainData(data); err != nil {
		return err
	}
	return nil
}

func (h *WearableHandler) updateWearableMinted(event wearables.WearableMintedEvent) error {
	data := wearables.WearablesOnChainData{
		Id:                 event.OnChainId,
		UpdatedAtTimestamp: event.Timestamp,
		UpdatedAtBlock:     event.BlockNumber,
	}
	if err := h.db.UpdateWearableOnMinted(data, event.Amount); err != nil {
		return err
	}
	activity := wearables.WearableActivities{
		OnChainId:    event.OnChainId,
		ActivityType: wearables.WearableActivityMinted,
		Timestamp:    event.Timestamp,
		BlockNumber:  event.BlockNumber,
		TxHash:       event.TxHash,
		Amount:       event.Amount,
		To:           event.To,
	}
	if err := h.db.AddWearableActivity(activity); err != nil {
		return err
	}
	return nil
}

func (h *WearableHandler) updateWearableBurnt(event wearables.WearableBurntEvent) error {
	data := wearables.WearablesOnChainData{
		Id:                 event.OnChainId,
		UpdatedAtTimestamp: event.Timestamp,
		UpdatedAtBlock:     event.BlockNumber,
	}
	if err := h.db.UpdateWearableOnBurnt(data, event.Amount); err != nil {
		return err
	}
	activity := wearables.WearableActivities{
		OnChainId:    event.OnChainId,
		ActivityType: wearables.WearableAcitivityBurnt,
		Timestamp:    event.Timestamp,
		BlockNumber:  event.BlockNumber,
		TxHash:       event.TxHash,
		Amount:       event.Amount,
		From:         event.From,
	}
	if err := h.db.AddWearableActivity(activity); err != nil {
		return err
	}
	return nil
}
