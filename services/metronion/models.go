package metronion

import (
	"math/big"

	"github.com/metrogalaxy-io/metrogalaxy-api/services"
)

type Metronion struct {
	Id                 uint64
	Name               string
	VersionId          uint64
	CreatedAtTimestamp uint64
	UpdatedAtTimestamp uint64
	CreatedAtBlock     uint64
	UpdatedAtBlock     uint64
	Owner              string
	LastPrice          float64
	CurrentPrice       float64
	TopBid             float64
	ListedBy           string
	Image              string
	Gender             string
	Type               string
	Star               uint64
	SpecialStar        uint64
	BaseStats          services.Stats
	SpecialStats       services.Stats
	Wearables          []Wearables
}

type MetronionOnChainData struct {
	Id                 uint64 `gorm:"primaryKey;autoIncrement:false"`
	Name               string
	VersionId          uint64
	CreatedAtTimestamp uint64
	UpdatedAtTimestamp uint64
	CreatedAtBlock     uint64
	UpdatedAtBlock     uint64
	Owner              string
	Wearables          string // list of metronion wearables id
}

type MetronionMetadata struct {
	Id           uint64
	Gender       string
	Type         string
	Image        string
	Star         uint64
	SpecialStar  uint64
	BaseStats    services.Stats
	SpecialStats services.Stats
	Wearables    []Wearables
}

type Wearables struct {
	WearableId string
	OnChainId  uint64
	Type       string
	Name       string
	Gender     string
	Rarity     string
	Changeable bool
	IsOrigin   bool
	IsRequired bool
	Image      string
	Stats      services.Stats
}

type MetronionFilterParams struct {
	Owner       string              `json:"owner"`
	Sort        string              `json:"sort"`
	MetronionID *uint64             `json:"id"`
	Gender      []string            `json:"gender"`
	Status      []string            `json:"status"`
	Stat        map[string][]uint64 `json:"stat"`
}

/**
 * Metronion Activity
 */
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

type MetronionActivity struct {
	ID           uint64 `gorm:"primaryKey"`
	MetronionID  uint64 `gorm:"uniqueIndex:metronion_id_activity_type_tx_hash;not null"`
	ActivityType string `gorm:"uniqueIndex:metronion_id_activity_type_tx_hash"`
	FromAccount  string
	ToAccount    string
	Price        float64 `gorm:"not null"`
	Timestamp    uint64
	BlockNumber  uint64
	TxHash       string `gorm:"uniqueIndex:metronion_id_activity_type_tx_hash"`
}

/**
 * Metronion Offers
 */
type MetronionOffers struct {
	ID          uint64  `gorm:"primaryKey"`
	MetronionID uint64  `gorm:"uniqueIndex:one_owner_one_metronion_offer;not null"`
	FromAccount string  `gorm:"uniqueIndex:one_owner_one_metronion_offer"`
	Price       float64 `gorm:"not null"`
	Timestamp   uint64
	BlockNumber uint64
	TxHash      string
}

/**
 * Metronion Listings
 */
type MetronionListing struct {
	ID          uint64 `gorm:"primaryKey"`
	MetronionID uint64 `gorm:"unique;not null"`
	FromAccount string
	Price       float64 `gorm:"not null"`
	Timestamp   uint64
	BlockNumber uint64
	TxHash      string
}

type MetronionPriceType string

const (
	MetronionPriceTypeLastPrice    MetronionPriceType = "last_price"
	MetronionPriceTypeCurrentPrice MetronionPriceType = "current_price"
	// MetronionPriceTypeTopBid       MetronionPriceType = "top_bid"
)

type MetronionPrice struct {
	ID           uint64 `gorm:"primaryKey"`
	MetronionID  uint64 `gorm:"unique"`
	LastPrice    float64
	CurrentPrice float64
}

/**
 * Version Config
 */

type MetronionVersionConfig struct {
	ID            uint64 `gorm:"primaryKey"`
	VersionID     uint64 `gorm:"unique"`
	StartingIndex uint64 `gorm:"not null"`
	Provenance    string
	MaxSupply     uint64
	CurrentSupply uint64 `gorm:"not null"`
}

type MetronionField string

const (
	MetronionFieldOwner     MetronionField = "owner"
	MetronionFieldName      MetronionField = "name"
	MetronionFieldWearables MetronionField = "wearables"
)

const (
	SortTypeNewest = "newest"
	SortTypeOldest = "oldest"
)

const (
	MetronionStatusForSale   = "for_sale"
	MetronionStatusHasOffers = "has_offers"
)

const (
	MetronionSortingLowestID        string = "lowest_id"
	MetronionSortingHighestID       string = "highest_id"
	MetronionSortingLowestPrice     string = "lowest_price"
	MetronionSortingHighestPrice    string = "highest_price"
	MetronionSortingRecentlyListed  string = "recently_listed"
	MetronionSortingRecentlySold    string = "recently_sold"
	MetronionSortingRecentlyOffered string = "recently_offered"
)

/**
 * Events
 */

type MetronionCreatedEvent struct {
	MetronionID        uint64
	VersionID          uint64
	Owner              string
	Name               string
	CreatedAtTimestamp uint64
	UpdatedAtTimestamp uint64
	CreatedAtBlock     uint64
	UpdatedAtBlock     uint64
	BaseActivity       services.OnChainBaseActivity
}

type MetronionTransferEvent struct {
	MetronionID uint64
	From        string
	To          string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type MetronionNameChangedEvent struct {
	MetronionID uint64
	NewName     string
	BlockNumber uint64
	Timestamp   uint64
}

type MetronionAccessoriesEquippedEvent struct {
	MetronionID  uint64
	AccessoryIds []uint64
	BlockNumber  uint64
	Timestamp    uint64
}

type MetronionAccessoriesUnequippedEvent struct {
	MetronionID  uint64
	AccessoryIds []uint64
	BlockNumber  uint64
	Timestamp    uint64
}

type MetronionListingEvent struct {
	MetronionID uint64
	Price       *big.Int
	Seller      string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type MetronionDelistEvent struct {
	MetronionID uint64
	Seller      string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type MetronionOfferEvent struct {
	MetronionID uint64
	Buyer       string
	Price       *big.Int
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type MetronionOfferCancelledEvent struct {
	MetronionID uint64
	Buyer       string
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type MetronionOfferTakenEvent struct {
	MetronionID uint64
	Buyer       string
	Seller      string
	Price       *big.Int
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}

type MetronionBuyEvent struct {
	MetronionID uint64
	Buyer       string
	Seller      string
	Price       *big.Int
	BlockNumber uint64
	Timestamp   uint64
	TxHash      string
}
