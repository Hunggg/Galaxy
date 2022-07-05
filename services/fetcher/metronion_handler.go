package fetcher

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"

	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/contracts"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/eventbus"
	"go.uber.org/zap"
)

type MetronionHandler struct {
	l                  *zap.SugaredLogger
	config             env.Config
	metronionDb        metronion.PersistentDb
	subscriptionKeeper *eventbus.SubscriptionKeeper
}

func NewMetronionHandler(metronionDb metronion.PersistentDb, subscriptionKeeper *eventbus.SubscriptionKeeper) *MetronionHandler {
	return &MetronionHandler{
		l:                  zap.S(),
		config:             env.GetConfig(),
		metronionDb:        metronionDb,
		subscriptionKeeper: subscriptionKeeper,
	}
}

func (h *MetronionHandler) Run(ctx context.Context, errCh chan<- error) {
	go h.SubscribeMetronionCreatedEvent(ctx, errCh)
	go h.SubscribeMetronionTransfer(ctx, errCh)
	go h.SubscribeNameChanged(ctx, errCh)
	go h.SubscribeAccessoriesEquipped(ctx, errCh)
	go h.SubscribeAccessoriesUnequipped(ctx, errCh)
}

func (h *MetronionHandler) SubscribeMetronionCreatedEvent(ctx context.Context, errCh chan<- error) {
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
			if err := h.SaveMetronion(value); err != nil {
				h.l.Errorw("error save metronion", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("save metronion created", "metronion_id", value.MetronionID, "owner", value.Owner, "created_at_timestamp", value.CreatedAtTimestamp)
		}
	}
}

func (h *MetronionHandler) SubscribeMetronionTransfer(ctx context.Context, errCh chan<- error) {
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

			data := metronion.MetronionOnChainData{
				Id:                 value.MetronionID,
				UpdatedAtTimestamp: value.Timestamp,
				UpdatedAtBlock:     value.BlockNumber,
				Owner:              value.To,
			}
			fields := []metronion.MetronionField{metronion.MetronionFieldOwner}

			if err := h.metronionDb.UpdateMetronionOnChainData(data, fields); err != nil {
				h.l.Errorw("error update metronion owner", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion owner", "metronion_id", value.MetronionID, "new_owner", value.To, "updated_at_timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionHandler) SubscribeNameChanged(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_NameChangedTopic.Hex(), internalErrCh)

	h.l.Infow("subscribe metronion name changed event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion name changed event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion name changed event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_NameChangedTopic.Hex(), internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionNameChangedEvent)
			if !ok {
				h.l.Errorw("error decode metronion name changed event", "message", message)
				errCh <- fmt.Errorf("error decode metronion name changed event")
				return
			}
			params := make(map[metronion.MetronionField]interface{})
			params[metronion.MetronionFieldName] = value.NewName

			data := metronion.MetronionOnChainData{
				Id:                 value.MetronionID,
				UpdatedAtTimestamp: value.Timestamp,
				UpdatedAtBlock:     value.BlockNumber,
				Name:               value.NewName,
			}
			fields := []metronion.MetronionField{metronion.MetronionFieldName}
			if err := h.metronionDb.UpdateMetronionOnChainData(data, fields); err != nil {
				h.l.Errorw("error update metronion accessories", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion accessories equipped", "metronion_id", value.MetronionID, "new_name", value.NewName, "updated_at_timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionHandler) SubscribeAccessoriesEquipped(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_AccessoriesEquippedTopic.Hex(), internalErrCh)

	h.l.Infow("subscribe metronion accessories equipped event", "sub_id", sub.Id)

	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion accessories equipped event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion accessories equipped event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_AccessoriesEquippedTopic.Hex(), internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionAccessoriesEquippedEvent)
			if !ok {
				h.l.Errorw("error decode metronion accessories equipped event", "message", message)
				errCh <- fmt.Errorf("error decode metronion accessories equipped event")
				return
			}
			if err := h.UpdateMetronionAccessories(value.MetronionID, value.BlockNumber, value.Timestamp, value.AccessoryIds, []uint64{}); err != nil {
				h.l.Errorw("error update metronion accessories", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion accessories equipped", "metronion_id", value.MetronionID, "accessory_ids", value.AccessoryIds, "updated_at_timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionHandler) SubscribeAccessoriesUnequipped(ctx context.Context, errCh chan<- error) {
	internalErrCh := make(chan error)
	sub := h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_AccessoriesUnequippedTopic.Hex(), internalErrCh)

	h.l.Infow("subscribe metronion accessories unequipped event", "sub_id", sub.Id)
	for {
		select {
		case <-ctx.Done():
			h.l.Debugw("unsubscribe metronion accessories unequipped event", "sub_id", sub.Id)
			h.subscriptionKeeper.Unsubscribe(sub.Id)
			return
		case err := <-internalErrCh:
			h.l.Errorw("subscribe metronion accessories unequipped event error. Retry subscribing...", "error", err)
			sub = h.subscriptionKeeper.Subscribe(contracts.MetronionNFT_AccessoriesUnequippedTopic.Hex(), internalErrCh)
		case message := <-sub.Receiver:
			value, ok := message.(metronion.MetronionAccessoriesUnequippedEvent)
			if !ok {
				h.l.Errorw("error decode metronion accessories unequipped event", "message", message)
				errCh <- fmt.Errorf("error decode metronion accessories unequipped event")
				return
			}
			if err := h.UpdateMetronionAccessories(value.MetronionID, value.BlockNumber, value.Timestamp, []uint64{}, value.AccessoryIds); err != nil {
				h.l.Errorw("error update metronion accessories", "error", err)
				errCh <- err
				return
			}
			h.l.Debugw("update metronion accessories unequipped", "metronion_id", value.MetronionID, "accessory_ids", value.AccessoryIds, "updated_at_timestamp", value.Timestamp)
		}
	}
}

func (h *MetronionHandler) SaveMetronion(event metronion.MetronionCreatedEvent) error {
	model := metronion.MetronionOnChainData{
		Id:                 event.MetronionID,
		CreatedAtTimestamp: event.CreatedAtTimestamp,
		UpdatedAtTimestamp: event.UpdatedAtTimestamp,
		CreatedAtBlock:     event.CreatedAtBlock,
		UpdatedAtBlock:     event.UpdatedAtBlock,
		VersionId:          event.VersionID,
		Owner:              event.Owner,
	}
	return h.metronionDb.SaveMetronionOnChainData(model)
}

func (h *MetronionHandler) UpdateMetronionAccessories(metronionID, updatedAtBlock, updatedAtTimestamp uint64, equippedAccessories, unequippedAccessories []uint64) error {
	metronionMetadata, err := h.metronionDb.GetMetronionById(metronionID)
	if err != nil {
		h.l.Debugw("error get metronion by id from db", "error", err)
		return err
	}

	var currentAccessoryIds = make([]string, len(metronionMetadata.Wearables))
	for i := range metronionMetadata.Wearables {
		currentAccessoryIds[i] = metronionMetadata.Wearables[i].WearableId
	}

	var mapAccessoryIds = make(map[string]bool)
	for _, item := range currentAccessoryIds {
		if item == "" {
			continue
		}
		mapAccessoryIds[item] = true
	}

	for _, item := range equippedAccessories {
		toStr := strconv.FormatUint(item, 10)
		if _, ok := mapAccessoryIds[toStr]; !ok {
			mapAccessoryIds[toStr] = true
		}
	}

	for _, item := range unequippedAccessories {
		toStr := strconv.FormatUint(item, 10)
		delete(mapAccessoryIds, toStr)
	}

	var newAccessoryIds = make([]string, 0)
	for key := range mapAccessoryIds {
		newAccessoryIds = append(newAccessoryIds, key)
	}

	params := make(map[metronion.MetronionField]interface{})
	params[metronion.MetronionFieldWearables] = strings.Join(newAccessoryIds, ",")

	data := metronion.MetronionOnChainData{
		Id:                 metronionID,
		UpdatedAtTimestamp: updatedAtTimestamp,
		UpdatedAtBlock:     updatedAtBlock,
		Wearables:          strings.Join(newAccessoryIds, ","),
	}
	fields := []metronion.MetronionField{metronion.MetronionFieldWearables}
	return h.metronionDb.UpdateMetronionOnChainData(data, fields)
}
