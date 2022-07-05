package cockroachdb

import (
	"context"
	"fmt"
	"strings"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/metrogalaxy-io/metrogalaxy-api/services"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	BatchSize                       = 500
	MetronionOffersIndex            = "one_owner_one_metronion_offer"
	MetronionActivitiesIndex        = "metronion_id_activity_type_tx_hash"
	MetronionListingTableName       = "metronion_listings"
	MetronionOnChainDataTableName   = "metronion_on_chain_data"
	MetronionPriceTableName         = "metronion_prices"
	MetronionVersionConfigTableName = "metronion_version_configs"
	MetronionOffersTableName        = "metronion_offers"
	MetronionActivitiesTableName    = "metronion_activities"
)

type CockroachDB struct {
	l  *zap.SugaredLogger
	db *gorm.DB
}

func NewCockroachDB(db *gorm.DB) (*CockroachDB, error) {
	l := zap.S()
	tables := make([]interface{}, 0)
	tables = append(tables, &MetronionMetadata{},
		&Wearables{},
		&metronion.MetronionOnChainData{},
		&metronion.MetronionVersionConfig{},
		&metronion.MetronionPrice{},
		&metronion.MetronionListing{})

	metronionOfferTable := &metronion.MetronionOffers{}
	metronionActivitiesTable := &metronion.MetronionActivity{}
	if !db.Migrator().HasTable(metronionOfferTable) {
		if err := db.Migrator().CreateTable(metronionOfferTable); err != nil {
			l.Errorw("error create metronion offers table", "error", err)
			return nil, err
		}
	}

	if !db.Migrator().HasIndex(metronionOfferTable, MetronionOffersIndex) {
		if err := db.Migrator().CreateIndex(metronionOfferTable, MetronionOffersIndex); err != nil {
			l.Errorw("error create metronion offers index", "error", err)
			return nil, err
		}
	}

	if !db.Migrator().HasTable(metronionActivitiesTable) {
		if err := db.Migrator().CreateTable(metronionActivitiesTable); err != nil {
			l.Errorw("error create metronion activities table", "error", err)
			return nil, err
		}
	}

	if !db.Migrator().HasIndex(metronionActivitiesTable, MetronionActivitiesIndex) {
		if err := db.Migrator().CreateIndex(metronionActivitiesTable, MetronionActivitiesIndex); err != nil {
			l.Errorw("error create metronion activities index", "error", err)
			return nil, err
		}
	}

	if err := db.AutoMigrate(tables...); err != nil {
		return nil, err
	}

	return &CockroachDB{
		l:  l,
		db: db,
	}, nil
}

func (c *CockroachDB) GetMetronionById(id uint64) (metronion.Metronion, error) {
	return getMetronionById(c.db, id)
}

func getWearables(db *gorm.DB, wearableId []string) ([]metronion.Wearables, error) {

	resultArr := []metronion.Wearables{}
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	for _, v := range wearableId {
		result := WearablesModelDB{}
		if err := db.WithContext(ctx).Raw(`select * from wearables w where wearable_id = ?`, v).Scan(&result).Error; err != nil {
			return []metronion.Wearables{}, err
		}
		stats := services.Stats{
			Creativity: result.Creativity,
			Charisma:   result.Charisma,
			Resolve:    result.Resolve,
			Fitness:    result.Fitness,
			Intellect:  result.Intellect,
		}
		wearableType := result.Type
		if wearableType == "" {
			if strings.Contains(strings.ToLower(v), "hair_long") {
				wearableType = string(services.MetronionWearableTypeHairLong)
			} else if strings.Contains(strings.ToLower(v), "head_skin") {
				wearableType = string(services.MetronionWearableTypeHeadSkin)
			}
		}

		elementWerable := metronion.Wearables{
			WearableId: v,
			Type:       wearableType,
			Name:       result.Name,
			Gender:     result.Gender,
			Rarity:     result.Rarity,
			Changeable: result.Changeable,
			IsOrigin:   result.IsOrigin,
			IsRequired: result.IsRequired,
			Image:      result.Image,
			Stats:      stats,
		}
		resultArr = append(resultArr, elementWerable)
	}
	return resultArr, nil
}

func getMetronionById(db *gorm.DB, id uint64) (metronion.Metronion, error) {
	var result Metronion
	var specialStats services.Stats
	var baseStats services.Stats
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := db.WithContext(ctx).Raw(`
	SELECT
		metronion_on_chain_data.id,
		metronion_on_chain_data.name,
		metronion_on_chain_data.version_id,
		metronion_on_chain_data.created_at_timestamp,
		metronion_on_chain_data.updated_at_timestamp,
		metronion_on_chain_data.created_at_block,
		metronion_on_chain_data.updated_at_block,
		metronion_on_chain_data.owner,
		metronion_prices.last_price,
		metronion_prices.current_price,
		top_bid,
		metronion_listings.from_account listed_by,
		metronion_metadata.image,
		metronion_metadata.gender,
		metronion_metadata.type,
		metronion_metadata.star,
		metronion_metadata.special_star,
		metronion_metadata.stats_creativity,
		metronion_metadata.stats_charisma,
		metronion_metadata.stats_resolve,
		metronion_metadata.stats_fitness,
		metronion_metadata.stats_intellect,
		metronion_metadata.special_stats_creativity,
		metronion_metadata.special_stats_charisma,
		metronion_metadata.special_stats_resolve,
		metronion_metadata.special_stats_fitness,
		metronion_metadata.special_stats_intellect,
		metronion_metadata.wearables
	FROM
		metronion_on_chain_data
		LEFT JOIN metronion_metadata ON metronion_metadata.id = metronion_on_chain_data.id
		LEFT JOIN metronion_prices ON metronion_prices.metronion_id = metronion_on_chain_data.id
		LEFT join metronion_listings on metronion_listings.metronion_id = metronion_on_chain_data.id 
		LEFT JOIN (
			SELECT
				max(price) top_bid,
				metronion_id
			FROM
				metronion_offers
			GROUP BY
				metronion_id) ON metronion_id = metronion_on_chain_data.id
	WHERE metronion_on_chain_data.id = ?`, id).Scan(&result).Error; err != nil {
		return metronion.Metronion{}, err
	}
	if result.CreatedAtTimestamp == 0 {
		return metronion.Metronion{}, fmt.Errorf("metronion not found")
	}

	baseStats.Creativity = result.StatsCreativity
	baseStats.Charisma = result.StatsCharisma
	baseStats.Resolve = result.StatsResolve
	baseStats.Fitness = result.StatsFitness
	baseStats.Intellect = result.StatsIntellect

	specialStats.Creativity = result.SpecialStatsCreativity
	specialStats.Charisma = result.SpecialStatsCharisma
	specialStats.Resolve = result.SpecialStatsResolve
	specialStats.Fitness = result.SpecialStatsFitness
	specialStats.Intellect = result.SpecialStatsIntellect
	wearablesArr := strings.Split(result.Wearables, ",")
	listWearables, _ := getWearables(db, wearablesArr)
	metronionResult := metronion.Metronion{
		Id:                 result.Id,
		Name:               result.Name,
		VersionId:          result.VersionId,
		CreatedAtTimestamp: result.CreatedAtTimestamp,
		UpdatedAtTimestamp: result.UpdatedAtTimestamp,
		CreatedAtBlock:     result.CreatedAtBlock,
		UpdatedAtBlock:     result.UpdatedAtBlock,
		Owner:              result.Owner,
		LastPrice:          result.LastPrice,
		CurrentPrice:       result.CurrentPrice,
		TopBid:             result.TopBid,
		ListedBy:           result.ListedBy,
		Image:              result.Image,
		Gender:             result.Gender,
		Type:               result.Type,
		Star:               result.Star,
		SpecialStar:        result.SpecialStar,
		BaseStats:          baseStats,
		SpecialStats:       specialStats,
		Wearables:          listWearables,
	}
	return metronionResult, nil
}

func (c *CockroachDB) GetListMetronions(offset int, limit int, filter metronion.MetronionFilterParams) ([]metronion.Metronion, uint64, error) {
	var result []Metronion
	var listMetronion []metronion.Metronion
	var specialStats services.Stats
	var baseStats services.Stats
	var count int64
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()

	var sortedBy = "id ASC"
	switch filter.Sort {
	case metronion.MetronionSortingLowestID:
		sortedBy = "id ASC"
	case metronion.MetronionSortingHighestID:
		sortedBy = "id DESC"
	case metronion.MetronionSortingLowestPrice:
		sortedBy = "CASE WHEN current_price > 0 then 1 ELSE 2 END, current_price ASC"
	case metronion.MetronionSortingHighestPrice:
		sortedBy = "current_price DESC"
	case metronion.MetronionSortingRecentlyListed:
	case metronion.MetronionSortingRecentlySold:
	case metronion.MetronionSortingRecentlyOffered:
	}

	query := c.db.WithContext(ctx).Table(MetronionOnChainDataTableName).Select(
		"metronion_on_chain_data.id," +
			"metronion_on_chain_data.name," +
			"metronion_on_chain_data.version_id," +
			"metronion_on_chain_data.created_at_timestamp," +
			"metronion_on_chain_data.updated_at_timestamp," +
			"metronion_on_chain_data.created_at_block," +
			"metronion_on_chain_data.updated_at_block," +
			"metronion_on_chain_data.owner," +
			"COALESCE(metronion_prices.last_price, 0) as last_price," +
			"COALESCE(metronion_prices.current_price, 0) as current_price," +
			"COALESCE(mo.price, 0) as top_bid," +
			"metronion_listings.from_account listed_by," +
			"metronion_metadata.image," +
			"metronion_metadata.gender," +
			"metronion_metadata.type," +
			"metronion_metadata.star," +
			"metronion_metadata.special_star," +
			"metronion_metadata.stats_creativity," +
			"metronion_metadata.stats_charisma," +
			"metronion_metadata.stats_resolve," +
			"metronion_metadata.stats_fitness," +
			"metronion_metadata.stats_intellect," +
			"metronion_metadata.special_stats_creativity," +
			"metronion_metadata.special_stats_charisma," +
			"metronion_metadata.special_stats_resolve," +
			"metronion_metadata.special_stats_fitness," +
			"metronion_metadata.special_stats_intellect," +
			"metronion_metadata.wearables").
		Joins("LEFT JOIN metronion_metadata ON metronion_metadata.id = metronion_on_chain_data.id ").
		Joins("LEFT JOIN metronion_prices ON metronion_prices.metronion_id = metronion_on_chain_data.id").
		Joins("LEFT join metronion_listings ON metronion_listings.metronion_id = metronion_on_chain_data.id").
		Joins("LEFT JOIN (SELECT MAX(price) as price, metronion_id FROM metronion_offers GROUP BY metronion_id) mo ON mo.metronion_id = metronion_on_chain_data.id")

	if filter.MetronionID != nil {
		metronionId := *filter.MetronionID
		listedMetronions := c.db.WithContext(ctx).Table(MetronionListingTableName).Where("metronion_id = ?", metronionId).Select("metronion_id")
		if filter.Owner != "" {
			listedMetronions.Where("from_account = ?", strings.ToLower(filter.Owner))
			query.Where("metronion_on_chain_data.owner = ?", strings.ToLower(filter.Owner))
		}
		query.Where("metronion_on_chain_data.id = ?", metronionId).Or("metronion_on_chain_data.id IN (?)", listedMetronions)
	} else {
		// filter gender
		if filter.Gender != nil && len(filter.Gender) > 0 {
			query.Where("metronion_metadata.gender IN (?)", filter.Gender)
		}
		// filter status
		if filter.Status != nil && len(filter.Status) > 0 {
			for _, status := range filter.Status {
				switch status {
				case metronion.MetronionStatusForSale:
					listedMetronions := c.db.WithContext(ctx).Table(MetronionListingTableName).Select("metronion_id")
					query.Where("metronion_on_chain_data.id IN (?)", listedMetronions)
				case metronion.MetronionStatusHasOffers:
					offeredMetronions := c.db.WithContext(ctx).Table(MetronionOffersTableName).Select("metronion_id").Group("metronion_id")
					query.Where("metronion_on_chain_data.id IN (?)", offeredMetronions)
				}
			}
		}

		// filter stats
		if filter.Stat != nil && len(filter.Stat) > 0 {
			var subQuery = c.db.WithContext(ctx).Table(`(SELECT
				id,
				CASE WHEN metronion_metadata.type = 'basic' THEN
					stats_creativity
				ELSE
					special_stats_creativity
				END AS creativity,
				CASE WHEN metronion_metadata.type = 'basic' THEN
					stats_charisma
				ELSE
					special_stats_charisma
				END AS charisma,
				CASE WHEN metronion_metadata.type = 'basic' THEN
					stats_resolve
				ELSE
					special_stats_resolve
				END AS resolve,
				CASE WHEN metronion_metadata.type = 'basic' THEN
					stats_fitness
				ELSE
					special_stats_fitness
				END AS fitness,
				CASE WHEN metronion_metadata.type = 'basic' THEN
					stats_intellect
				ELSE
					special_stats_intellect
				END AS intellect
			FROM
				metronion_metadata)`).Select("id")

			for stat, values := range filter.Stat {
				if len(values) != 2 {
					return nil, 0, fmt.Errorf(fmt.Sprintf("invalid value range for stat %s", stat))
				}
				subQuery.Where(fmt.Sprintf("%s >= ? AND %s <= ?", stat, stat), values[0], values[1])
			}
			subQuery.Select("id")

			query.Where("metronion_on_chain_data.id IN (?)", subQuery)
		}

		// filter owner
		if filter.Owner != "" {
			owner := strings.ToLower(filter.Owner)
			listedMetronions := c.db.WithContext(ctx).Table(MetronionListingTableName).Where("from_account = ?", owner).Select("metronion_id")
			query.Where(c.db.Where("metronion_on_chain_data.owner = ?", owner).Or("metronion_on_chain_data.id IN (?)", listedMetronions))
		}
	}

	if err := query.Count(&count).Error; err != nil {
		c.l.Errorw("error get list metronions count", "error", err)
		return nil, 0, err
	}

	if err :=
		query.Offset(offset * limit).
			Limit(limit).Order(sortedBy).Scan(&result).Error; err != nil {
		c.l.Errorw("error query get list metronions", "error", err)
		return nil, 0, err
	}
	for _, v := range result {
		baseStats.Creativity = v.StatsCreativity
		baseStats.Charisma = v.StatsCharisma
		baseStats.Resolve = v.StatsResolve
		baseStats.Fitness = v.StatsFitness
		baseStats.Intellect = v.StatsIntellect

		specialStats.Creativity = v.SpecialStatsCreativity
		specialStats.Charisma = v.SpecialStatsCharisma
		specialStats.Resolve = v.SpecialStatsResolve
		specialStats.Fitness = v.SpecialStatsFitness
		specialStats.Intellect = v.SpecialStatsIntellect
		wearablesArr := strings.Split(v.Wearables, ",")
		listWearables, _ := getWearables(c.db, wearablesArr)
		metronionResult := metronion.Metronion{
			Id:                 v.Id,
			Name:               v.Name,
			VersionId:          v.VersionId,
			CreatedAtTimestamp: v.CreatedAtTimestamp,
			UpdatedAtTimestamp: v.UpdatedAtTimestamp,
			CreatedAtBlock:     v.CreatedAtBlock,
			UpdatedAtBlock:     v.UpdatedAtBlock,
			Owner:              v.Owner,
			LastPrice:          v.LastPrice,
			CurrentPrice:       v.CurrentPrice,
			TopBid:             v.TopBid,
			ListedBy:           v.ListedBy,
			Image:              v.Image,
			Gender:             v.Gender,
			Type:               v.Type,
			Star:               v.Star,
			SpecialStar:        v.SpecialStar,
			BaseStats:          baseStats,
			SpecialStats:       specialStats,
			Wearables:          listWearables,
		}
		listMetronion = append(listMetronion, metronionResult)
	}
	return listMetronion, uint64(count), nil
}

func saveMetronion(db *gorm.DB, metronion metronion.MetronionOnChainData) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&metronion).Error
}
func saveBatchMetronion(db *gorm.DB, data []metronion.MetronionOnChainData) error {
	dbSession := db.Session(&gorm.Session{CreateBatchSize: BatchSize})
	return dbSession.Clauses(clause.OnConflict{DoNothing: true}).Create(&data).Error
}

func (c *CockroachDB) SaveMetronionOnChainData(data metronion.MetronionOnChainData) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveMetronion(tx, data)
		},
	); err != nil {
		c.l.Errorw("error save metronion", "error", err)
		return err
	}
	return nil
}

func (c *CockroachDB) SaveBatchMetronionOnChainData(listData []metronion.MetronionOnChainData) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveBatchMetronion(tx, listData)
		},
	); err != nil {
		c.l.Errorw("error save batch metronion", "error", err)
		return err
	}
	return nil
}

func (c *CockroachDB) UpdateMetronionOnChainData(data metronion.MetronionOnChainData, fields []metronion.MetronionField) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()

	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return updateMetronionOnChainData(tx, data, fields)
		},
	); err != nil {
		c.l.Errorw("error update metronion onchain data", "error", err)
		return err
	}
	return nil
}

func updateMetronionOnChainData(db *gorm.DB, data metronion.MetronionOnChainData, fields []metronion.MetronionField) error {
	insertData := data
	insertData.Owner = strings.ToLower(insertData.Owner)
	updatedFields := make([]string, 0)
	updatedFields = append(updatedFields, "updated_at_block", "updated_at_timestamp")
	for _, f := range fields {
		updatedFields = append(updatedFields, string(f))
	}

	return db.Model(&metronion.MetronionOnChainData{}).Where("id = ?", data.Id).Select(updatedFields).UpdateColumns(insertData).Error
}

func (c *CockroachDB) GetMetronionOnChainData(id uint64) (metronion.MetronionOnChainData, error) {
	var result metronion.MetronionOnChainData
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := c.db.WithContext(ctx).Where("id", id).Find(&result); err != nil {
		c.l.Errorw("error get information of MetronionOnchainData", "error", err, "metronion_id", id)
		return metronion.MetronionOnChainData{}, err.Error
	}
	return result, nil
}

/**
 * Metronion Metadata
 */
func (c *CockroachDB) SaveBatchMetronionMetadata(data []metronion.MetronionMetadata) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveBatchMetronionMetadata(tx, data)
		},
	); err != nil {
		c.l.Errorw("error save batch metronion metadata", "error", err)
		return err
	}
	return nil
}

func saveBatchMetronionMetadata(db *gorm.DB, data []metronion.MetronionMetadata) error {
	dbSession := db.Session(&gorm.Session{CreateBatchSize: BatchSize})

	var records = make([]MetronionMetadata, len(data))
	for i := range data {
		records[i] = ConvertMetronionMetadata(data[i])
	}

	return dbSession.Create(&records).Error
}

/**
 * Wearable Metadata
 */
func (c *CockroachDB) SaveBatchWearableMetadata(data []metronion.Wearables) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveBatchWearableMetadata(tx, data)
		},
	); err != nil {
		c.l.Errorw("error save batch wearable metadata", "error", err)
		return err
	}
	return nil
}

func saveBatchWearableMetadata(db *gorm.DB, data []metronion.Wearables) error {
	dbSession := db.Session(&gorm.Session{CreateBatchSize: BatchSize})

	var records = make([]Wearables, len(data))
	for i := range data {
		records[i] = ConvertWearables(data[i])
	}

	return dbSession.Create(&records).Error
}

/**
 * Metronion Activities
 */
func (c *CockroachDB) SaveMetronionActivity(activity metronion.MetronionActivity) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveMetronionActivity(tx, activity)
		},
	); err != nil {
		c.l.Errorw("error save metronion activity", "error", err)
		return err
	}
	return nil
}

func saveMetronionActivity(db *gorm.DB, activity metronion.MetronionActivity) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&activity).Error
}

func (c *CockroachDB) GetMetronionActivities(metronionID uint64, sortType string) ([]metronion.MetronionActivity, error) {
	var result = make([]metronion.MetronionActivity, 0)
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()

	var sortByTimestamp string = "DESC"
	if sortType == metronion.SortTypeOldest {
		sortByTimestamp = "ASC"
	}

	if err := c.db.WithContext(ctx).Where("metronion_id = ?", metronionID).Order(fmt.Sprintf("timestamp %s", sortByTimestamp)).Find(&result).Error; err != nil {
		c.l.Errorw("error get metronion activities", "error", err, "metronion_id", metronionID)
		return nil, err
	}

	return result, nil
}

/**
 * Metronion Offers
 */

// UpdateMetronionOffers insert or delete metronion offers by MetronionID and FromAccount
func (c *CockroachDB) UpdateMetronionOffers(offer metronion.MetronionOffers, isAdd bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return updateMetronionOffers(tx, offer, isAdd)
		},
	); err != nil {
		c.l.Errorw("error update metronion offer", "error", err)
		return err
	}
	return nil
}

func updateMetronionOffers(db *gorm.DB, offer metronion.MetronionOffers, isAdd bool) error {
	if !isAdd {
		return db.Where("metronion_id = ? AND from_account = ?", offer.MetronionID, offer.FromAccount).Delete(&offer).Error
	}
	offer.FromAccount = strings.ToLower(offer.FromAccount)
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&offer).Error
}

// GetMetronionOffers get metronion offers in higher price descending order
func (c *CockroachDB) GetMetronionOffers(metronionID uint64, sortType string) ([]metronion.MetronionOffers, error) {
	var result = make([]metronion.MetronionOffers, 0)
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()

	var order string = "price DESC"
	if sortType == metronion.SortTypeOldest {
		order = "timestamp ASC"
	} else if sortType == metronion.SortTypeNewest {
		order = "timestamp DESC"
	}
	if err := c.db.WithContext(ctx).Where("metronion_id = ?", metronionID).Order(order).Find(&result).Error; err != nil {
		c.l.Errorw("error get metronion offers", "error", err, "metronion_id", metronionID)
		return nil, err
	}

	return result, nil
}

/**
 * Metronion Listing
 */

// UpdateMetronionListing insert or delete metronion listing, if the listing of specific MetronionID is existed, then don't update anything
func (c *CockroachDB) UpdateMetronionListing(listing metronion.MetronionListing, isAdd bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return updateMetronionListing(tx, listing, isAdd)
		},
	); err != nil {
		c.l.Errorw("error update metronion listing", "error", err)
		return err
	}
	return nil
}

func updateMetronionListing(db *gorm.DB, listing metronion.MetronionListing, isAdd bool) error {
	if !isAdd {
		return db.Where("metronion_id = ?", listing.MetronionID).Delete(&listing).Error
	}
	listing.FromAccount = strings.ToLower(listing.FromAccount)
	return db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&listing).Error
}

func (c *CockroachDB) GetMetronionListing(metronionID uint64) (metronion.MetronionListing, error) {
	var result metronion.MetronionListing
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := c.db.WithContext(ctx).Where("metronion_id = ?", metronionID).Take(&result).Error; err != nil {
		c.l.Errorw("error get metronion listing", "error", err, "metronion_id", metronionID)
		return metronion.MetronionListing{}, err
	}

	return result, nil
}

/**
 * Metronion Price
 */
// UpdateMetronionPrice upsert metronion price by MetronionID, update price by price type
func (c *CockroachDB) UpdateMetronionPrice(metronionID uint64, priceType metronion.MetronionPriceType, value float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return updateMetronionPrice(tx, metronionID, priceType, value)
		},
	); err != nil {
		c.l.Errorw("error update metronion price", "error", err, "metronion_id", metronionID, "price_type", priceType, "value", value)
		return err
	}
	return nil
}

// updateMetronionPrice upsert metronion price
func updateMetronionPrice(db *gorm.DB, metronionID uint64, priceType metronion.MetronionPriceType, value float64) error {
	var updatedColumns = make([]string, 0)
	updatedColumns = append(updatedColumns, string(priceType))

	var record = &metronion.MetronionPrice{
		MetronionID: metronionID,
	}
	if priceType == metronion.MetronionPriceTypeCurrentPrice {
		record.CurrentPrice = value
	} else if priceType == metronion.MetronionPriceTypeLastPrice {
		record.LastPrice = value
	}

	return db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "metronion_id"},
		},
		DoUpdates: clause.AssignmentColumns(updatedColumns),
	}).Create(record).Error
}

func (c *CockroachDB) GetMetronionPrice(metronionID uint64) (metronion.MetronionPrice, error) {
	var result metronion.MetronionPrice
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := c.db.WithContext(ctx).Where("metronion_id = ?", metronionID).Take(&result).Error; err != nil {
		c.l.Errorw("error get metronion price", "error", err, "metronion_id", metronionID)
		return metronion.MetronionPrice{}, err
	}

	return result, nil
}

/**
 * Version Config
 */
func (c *CockroachDB) SaveMetronionVersionConfig(version metronion.MetronionVersionConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveMetronionVersionConfig(tx, version)
		},
	); err != nil {
		c.l.Errorw("error save metronion version config", "error", err)
		return err
	}
	return nil
}

func saveMetronionVersionConfig(db *gorm.DB, version metronion.MetronionVersionConfig) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&version).Error
}

func (c *CockroachDB) GetMetronionVersionConfig(versionID uint64) (metronion.MetronionVersionConfig, error) {
	var result metronion.MetronionVersionConfig
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := c.db.WithContext(ctx).Where("version_id = ?", versionID).Take(&result).Error; err != nil {
		c.l.Errorw("error get metronion version config", "error", err)
		return metronion.MetronionVersionConfig{}, err
	}

	return result, nil
}