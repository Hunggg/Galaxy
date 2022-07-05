package wearables

type PersistentDb interface {
	SaveWearableOnChainData(wearable WearablesOnChainData) error
	UpdateWearableOnMinted(wearable WearablesOnChainData, amount uint64) error
	UpdateWearableOnBurnt(wearable WearablesOnChainData, amount uint64) error
	AddWearableActivity(activity WearableActivities) error

	/*
	 * Wearable Listings
	 */
	UpdateWearableListing(listing WearableListing, isAdd bool) error

	/**
	 * Wearable Offers
	 */
	UpdateWearableOffer(offer WearableOffers, isAdd bool) error

	GetListWearablesActivity(offset int, limit int) ([]WearableActivities, uint64, error)
	GetListWearablesInformation(offset int, limit int, filter WearablesFilterParams) ([]WearableInfo, uint64, error)
	GetListWearablesListing(offset int, limit int) ([]WearableListing, uint64, error)
	GetListWearablesOffer(offset int, limit int) ([]WearableOffers, uint64, error)

	GetWearablesListing(id uint64) (WearableListing, error)
	GetWearablesOffer(id uint64) (WearableOffers, error)
	GetWearablesInformation(id uint64) (WearableInfo, error)
}

type WearableField string
