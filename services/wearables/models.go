package wearables

import (
	"math/big"

	"github.com/metrogalaxy-io/metrogalaxy-api/services"
)

const (
	ActivityTypeMint           = "mint"
	ActivityTypeTransfer       = "transfer"
	ActivityTypeList           = "list"
	ActivityTypeDelist         = "delist"
	ActivityTypeOffer          = "offer"
	ActivityTypeOfferCancelled = "offer_cancelled"
	ActivityTypeOfferTaken     = "offer_taken"
	ActivityTypeSale           = "sale"
)

type WearablesFilterParams struct {
	Rarity string   `json:"rarity"`
	Sort   string   `json:"sort"`
	Gender []string `json:"gender"`
	Type   []string `json:"type"`
}

type WearableInfo struct {
	WearableId         string
	Type               string 
	Name               string // defined on chain
	Gender             string
	Rarity             string
	Changeable         bool
	IsOrigin           bool
	IsRequired         bool
	Image              string
	Stats              services.Stats
	MaxSupply          uint64 // defined on chain
	CurrentSupply      uint64 // defined on chain
	CreatedAtTimestamp uint64 // defined on chain
	UpdatedAtTimestamp uint64 // defined on chain
	CreatedAtBlock     uint64 // defined on chain
	UpdatedAtBlock     uint64 // defined on chain
	OnChainId          uint64
}

type WearableInfoModel struct {
	WearableId         string
	OnChainId          uint64
	Name               string // defined on chain
	Gender             string
	Changeable         bool
	IsOrigin           bool
	IsRequired         bool
	Image              string
	Creativity         uint64
	Charisma           uint64
	Resolve            uint64
	Fitness            uint64
	Intellect          uint64
	MaxSupply          uint64 // defined on chain
	CurrentSupply      uint64 // defined on chain
	Type               string // defined on chain
	Rarity             string // defined on chain
	CreatedAtTimestamp uint64 // defined on chain
	UpdatedAtTimestamp uint64 // defined on chain
	CreatedAtBlock     uint64 // defined on chain
	UpdatedAtBlock     uint64 // defined on chain
}

type WearablesOnChainData struct {
	Id                 uint64 `gorm:"primaryKey;autoIncrement:false"`
	Name               string
	MaxSupply          uint64
	CurrentSupply      uint64
	Type               string
	Rarity             string
	CreatedAtTimestamp uint64
	UpdatedAtTimestamp uint64
	CreatedAtBlock     uint64
	UpdatedAtBlock     uint64
}

type WearableActivities struct {
	ID           uint64 `gorm:"primaryKey"`
	OnChainId    uint64 `gorm:"uniqueIndex:wearable_id_activity_type_tx_hash;not null"`
	ActivityType string `gorm:"uniqueIndex:wearable_id_activity_type_tx_hash"`
	Timestamp    uint64
	BlockNumber  uint64
	TxHash       string `gorm:"uniqueIndex:wearable_id_activity_type_tx_hash"`
	Amount       uint64
	Price        uint64
	From         string
	To           string
}

const (
	WearableActivityMinted = "mint"
	WearableAcitivityBurnt = "burnt"
)

type WearableListing struct {
	ID          uint64  `gorm:"primaryKey"`
	OnChainId   uint64  `gorm:"uniqueIndex:one_owner_one_wearable_listing;not null"`
	FromAccount string  `gorm:"uniqueIndex:one_owner_one_wearable_listing"`
	Price       float64 `gorm:"not null"`
	Amount      uint64
	Timestamp   uint64
	BlockNumber uint64
	TxHash      string
}

type WearableOffers struct {
	ID          uint64  `gorm:"primaryKey"`
	OnChainId   uint64  `gorm:"uniqueIndex:one_owner_one_wearable_offer;not null"`
	FromAccount string  `gorm:"uniqueIndex:one_owner_one_wearable_offer"`
	Price       float64 `gorm:"not null"`
	Amount      uint64
	Timestamp   uint64
	BlockNumber uint64
	TxHash      string
}

/**
 * On Chain Event
 */
type WearableCreatedEvent struct {
	OnChainId   uint64
	Name        string
	MaxSupply   uint64
	Type        string
	Rarity      string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableMintedEvent struct {
	OnChainId   uint64
	Amount      uint64
	To          string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableTransferEvent struct {
	OnChainId   uint64
	Amount      uint64
	From        string
	To          string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableBatchTransferEvent struct {
	OnChainIds  []uint64
	Amounts     []uint64
	From        string
	To          string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableBurntEvent struct {
	OnChainId   uint64
	Amount      uint64
	From        string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableStoredEvent struct {
	OnChainIds  []uint64
	From        string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearbleReturnedEvent struct {
	OnChainIds  []uint64
	Amount      uint64
	From        string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableListingEvent struct {
	OnChainId   uint64
	Price       *big.Int
	Amount      *big.Int
	Seller      string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableDelistEvent struct {
	OnChainId   uint64
	Seller      string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableOfferEvent struct {
	OnChainId   uint64
	Price       *big.Int
	Amount      *big.Int
	Buyer       string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableOfferCancelledEvent struct {
	OnChainId   uint64
	Buyer       string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableOfferTakenEvent struct {
	OnChainId   uint64
	Price       *big.Int
	Amount      *big.Int
	Buyer       string
	Seller      string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type WearableBoughtEvent struct {
	OnChainId   uint64
	Price       *big.Int
	Amount      *big.Int
	Buyer       string
	Seller      string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

const (
	WearablesSortingLowestCurrentSupply  string = "lowest_current_supply"
	WearablesSortingHighestCurrentSupply string = "highest_current_supply"
	WearablesSortingLowestMaxSupply      string = "lowest_max_supply"
	WearablesSortingHighestMaxSupply     string = "highest_max_supply"
)
