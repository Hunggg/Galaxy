package metronion

type PersistentDb interface {
	/**
	 * Metronions
	 */
	// SaveMetronion(metronion Metronion) error
	// SaveBatchMetronion(lists []Metronion) error
	// UpdateMetronion(id uint64, updatedAtBlock uint64, updatedAtTimestamp uint64, params map[MetronionField]interface{}) error
	GetMetronionById(id uint64) (Metronion, error)
	GetListMetronions(offset int, limit int, filter MetronionFilterParams) ([]Metronion, uint64, error)

	SaveMetronionOnChainData(data MetronionOnChainData) error
	SaveBatchMetronionOnChainData(listData []MetronionOnChainData) error
	UpdateMetronionOnChainData(data MetronionOnChainData, fields []MetronionField) error
	GetMetronionOnChainData(id uint64) (MetronionOnChainData, error)

	/**
	 * Metronion Metadata
	 */
	SaveBatchMetronionMetadata(data []MetronionMetadata) error

	/**
	 * Wearable Metadata
	 */
	SaveBatchWearableMetadata(data []Wearables) error

	/**
	 * Metronion Activities
	 */
	SaveMetronionActivity(activity MetronionActivity) error
	GetMetronionActivities(metronionID uint64, sortType string) ([]MetronionActivity, error)

	/**
	 * Metronion Offers
	 */
	UpdateMetronionOffers(offer MetronionOffers, isAdd bool) error
	GetMetronionOffers(metronionID uint64, sortType string) ([]MetronionOffers, error)

	/**
	 * Metronion Listing
	 */
	UpdateMetronionListing(listing MetronionListing, isAdd bool) error
	GetMetronionListing(metronionID uint64) (MetronionListing, error)

	/**
	 * Metronion Price
	 */
	UpdateMetronionPrice(metronionID uint64, priceType MetronionPriceType, value float64) error
	GetMetronionPrice(metronionID uint64) (MetronionPrice, error)

	/**
	 * Version Config
	 */
	SaveMetronionVersionConfig(version MetronionVersionConfig) error
	GetMetronionVersionConfig(versionID uint64) (MetronionVersionConfig, error)
}