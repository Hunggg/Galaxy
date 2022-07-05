package contracts

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	etherCommon "github.com/ethereum/go-ethereum/common"
)

var (
	MetronionNFT_TransferTopic              etherCommon.Hash
	MetronionNFT_MetronionCreatedTopic      etherCommon.Hash
	MetronionNFT_AccessoriesEquippedTopic   etherCommon.Hash
	MetronionNFT_AccessoriesUnequippedTopic etherCommon.Hash
	MetronionNFT_NameChangedTopic           etherCommon.Hash
	MetronionNFT_UpdateBaseURITopic         etherCommon.Hash

	MetronionSale_PrivateBoughtTopic etherCommon.Hash
	MetronionSale_PublicBoughtTopic  etherCommon.Hash

	MetronionAccessories_TransferSingleTopic   etherCommon.Hash
	MetronionAccessories_TransferBatchTopic    etherCommon.Hash
	MetronionAccessories_AccessoryCreatedTopic etherCommon.Hash
	MetronionAccessories_AccessoryMintTopic    etherCommon.Hash
	MetronionAccessories_AccessoryBurntTopic   etherCommon.Hash
	MetronionAccessories_AccessoryStoreTopic   etherCommon.Hash
	MetronionAccessories_AccessoryReturnTopic  etherCommon.Hash

	Marketplace_AssetListedTopic         etherCommon.Hash
	Marketplace_AssetDelistedTopic       etherCommon.Hash
	Marketplace_AssetBoughtTopic         etherCommon.Hash
	Marketplace_AssetOfferedTopic        etherCommon.Hash
	Marketplace_AssetOfferCancelledTopic etherCommon.Hash
	Marketplace_AssetOfferTakenTopic     etherCommon.Hash

	MetronionListingEventId        = "metronion_listing"
	MetronionDelistEventId         = "metronion_delist"
	MetronionOfferEventId          = "metronion_offer"
	MetronionOfferCancelledEventId = "metronion_cancelled"
	MetronionOfferTakenEventId     = "metronion_taken"
	MetronionBuyEventId            = "metronion_buy"
	WearableListingEventId         = "wearable_listing"
	WearableDelistEventId          = "wearable_delist"
	WearableOfferEventId           = "wearable_offer"
	WearableOfferCancelledEventId  = "wearable_offer_cancelled"
	WearableOfferTakenEventId      = "wearable_offer_taken"
	WearableBoughtEventId          = "wearable_bought"
)

func init() {
	MetronionNFT_TransferTopic = getContractTopic(MetronionNFTABI, "Transfer")
	MetronionNFT_MetronionCreatedTopic = getContractTopic(MetronionNFTABI, "MetronionCreated")
	MetronionNFT_AccessoriesEquippedTopic = getContractTopic(MetronionNFTABI, "AccessoriesEquipped")
	MetronionNFT_AccessoriesUnequippedTopic = getContractTopic(MetronionNFTABI, "AccessoriesUnequipped")
	MetronionNFT_NameChangedTopic = getContractTopic(MetronionNFTABI, "NameChanged")
	MetronionNFT_UpdateBaseURITopic = getContractTopic(MetronionNFTABI, "UpdateBaseURI")

	MetronionSale_PrivateBoughtTopic = getContractTopic(MetronionSaleABI, "PrivateBought")
	MetronionSale_PublicBoughtTopic = getContractTopic(MetronionSaleABI, "PublicBought")

	MetronionAccessories_TransferSingleTopic = getContractTopic(MetronionAccessoriesABI, "TransferSingle")
	MetronionAccessories_TransferBatchTopic = getContractTopic(MetronionAccessoriesABI, "TransferBatch")
	MetronionAccessories_AccessoryCreatedTopic = getContractTopic(MetronionAccessoriesABI, "AccessoryCreated")
	MetronionAccessories_AccessoryMintTopic = getContractTopic(MetronionAccessoriesABI, "AccessoryMint")
	MetronionAccessories_AccessoryBurntTopic = getContractTopic(MetronionAccessoriesABI, "AccessoryBurnt")
	MetronionAccessories_AccessoryStoreTopic = getContractTopic(MetronionAccessoriesABI, "AccessoryStore")
	MetronionAccessories_AccessoryReturnTopic = getContractTopic(MetronionAccessoriesABI, "AccessoryReturn")

	Marketplace_AssetListedTopic = getContractTopic(MetroGalaxyMarketplaceABI, "AssetListed")
	Marketplace_AssetDelistedTopic = getContractTopic(MetroGalaxyMarketplaceABI, "AssetDelisted")
	Marketplace_AssetBoughtTopic = getContractTopic(MetroGalaxyMarketplaceABI, "AssetBought")
	Marketplace_AssetOfferedTopic = getContractTopic(MetroGalaxyMarketplaceABI, "AssetOffered")
	Marketplace_AssetOfferCancelledTopic = getContractTopic(MetroGalaxyMarketplaceABI, "AssetOfferCancelled")
	Marketplace_AssetOfferTakenTopic = getContractTopic(MetroGalaxyMarketplaceABI, "AssetOfferTaken")
}

// getContractTopic get contract topics hash for event
func getContractTopic(contractABI string, eventName string) etherCommon.Hash {
	contractAbi, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Panic(err)
		return etherCommon.Hash{}
	}

	v, ok := contractAbi.Events[eventName]
	if !ok {
		log.Panicf("getContractTopic: event %v not found", eventName)
		return etherCommon.Hash{}
	}
	return v.ID
}
