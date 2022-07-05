package cockroachdb

import (
	"context"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/metrogalaxy-io/metrogalaxy-api/services"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/wearables"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	WearableActivitiesIndex = "wearable_id_activity_type_tx_hash"
	WearableListingIndex    = "one_owner_one_wearable_listing"
	WearableOffersIndex     = "one_owner_one_wearable_offer"
)

type CockroachDB struct {
	l  *zap.SugaredLogger
	db *gorm.DB
}

func NewCockroachDB(db *gorm.DB) (*CockroachDB, error) {
	l := zap.S()
	tables := make([]interface{}, 0)
	tables = append(tables, &wearables.WearablesOnChainData{})

	wearableActivitiesTable := &wearables.WearableActivities{}
	wearableListingTable := &wearables.WearableListing{}
	wearableOffersTable := &wearables.WearableOffers{}

	if !db.Migrator().HasTable(wearableActivitiesTable) {
		if err := db.Migrator().CreateTable(wearableActivitiesTable); err != nil {
			l.Errorw("error create wearable activities table", "error", err)
			return nil, err
		}
	}

	if !db.Migrator().HasIndex(wearableActivitiesTable, WearableActivitiesIndex) {
		if err := db.Migrator().CreateIndex(wearableActivitiesTable, WearableActivitiesIndex); err != nil {
			l.Errorw("error create wearable activities index", "error", err)
			return nil, err
		}
	}

	if !db.Migrator().HasTable(wearableListingTable) {
		if err := db.Migrator().CreateTable(wearableListingTable); err != nil {
			l.Errorw("error create wearable listing table", "error", err)
			return nil, err
		}
	}

	if !db.Migrator().HasIndex(wearableListingTable, WearableListingIndex) {
		if err := db.Migrator().CreateIndex(wearableListingTable, WearableListingIndex); err != nil {
			l.Errorw("error create wearable listing index", "error", err)
			return nil, err
		}
	}

	if !db.Migrator().HasTable(wearableOffersTable) {
		if err := db.Migrator().CreateTable(wearableOffersTable); err != nil {
			l.Errorw("error create wearable offers table", "error", err)
			return nil, err
		}
	}

	if !db.Migrator().HasIndex(wearableOffersTable, WearableOffersIndex) {
		if err := db.Migrator().CreateIndex(wearableOffersTable, WearableOffersIndex); err != nil {
			l.Errorw("error create wearable offers index", "error", err)
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

func (c *CockroachDB) SaveWearableOnChainData(data wearables.WearablesOnChainData) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return saveWearableOnChainData(tx, data)
		},
	); err != nil {
		c.l.Errorw("error save wearable on chain data", "error", err)
		return err
	}
	return nil
}

func saveWearableOnChainData(db *gorm.DB, data wearables.WearablesOnChainData) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&data).Error
}

func (c *CockroachDB) UpdateWearableOnMinted(data wearables.WearablesOnChainData, amount uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()

	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return updateWearableOnMinted(tx, data, amount)
		},
	); err != nil {
		c.l.Errorw("error update wearable on minted", "error", err)
		return err
	}
	return nil
}

func updateWearableOnMinted(db *gorm.DB, data wearables.WearablesOnChainData, amount uint64) error {
	return db.Model(&wearables.WearablesOnChainData{}).Where("id = ?", data.Id).Updates(map[string]interface{}{
		"updated_at_timestamp": data.UpdatedAtTimestamp,
		"updated_at_block":     data.UpdatedAtBlock,
		"current_supply":       gorm.Expr("current_supply + ?", amount),
	}).Error
}

func (c *CockroachDB) UpdateWearableOnBurnt(data wearables.WearablesOnChainData, amount uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()

	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return updateWearableOnBurnt(tx, data, amount)
		},
	); err != nil {
		c.l.Errorw("error update wearable on burnt", "error", err)
		return err
	}
	return nil
}

func updateWearableOnBurnt(db *gorm.DB, data wearables.WearablesOnChainData, amount uint64) error {
	return db.Model(&wearables.WearablesOnChainData{}).Where("on_chain_id = ?", data.Id).Updates(map[string]interface{}{
		"updated_at_timestamp": data.UpdatedAtTimestamp,
		"updated_at_block":     data.UpdatedAtBlock,
		"current_supply":       gorm.Expr("current_supply - ?", amount),
		"max_supply":           gorm.Expr("max_supply - ?", amount),
	}).Error
}

func (c *CockroachDB) AddWearableActivity(activity wearables.WearableActivities) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()

	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&activity).Error
		},
	); err != nil {
		c.l.Errorw("error add wearable activity", "error", err)
		return err
	}
	return nil
}

func (c *CockroachDB) UpdateWearableListing(listing wearables.WearableListing, isAdd bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()

	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return updateWearablesListing(tx, listing, isAdd)
		},
	); err != nil {
		c.l.Errorw("error update wearables listing", "error", err)
		return err
	}
	return nil
}

func updateWearablesListing(db *gorm.DB, listing wearables.WearableListing, isAdd bool) error {
	if !isAdd {
		return db.Where("on_chain_id = ? AND from_account = ?", listing.OnChainId, listing.FromAccount).Delete(&listing).Error
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&listing).Error
}

func (c *CockroachDB) UpdateWearableOffer(offer wearables.WearableOffers, isAdd bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()

	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return updateWearablesOffer(tx, offer, isAdd)
		},
	); err != nil {
		c.l.Errorw("error update wearables offer", "error", err)
		return err
	}
	return nil
}

func updateWearablesOffer(db *gorm.DB, offer wearables.WearableOffers, isAdd bool) error {
	if !isAdd {
		return db.Where("on_chain_id = ? AND from_account = ?", offer.OnChainId, offer.FromAccount).Delete(&offer).Error
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&offer).Error
}

func (c *CockroachDB) getWearablesOnChainDataById(id uint64) (wearables.WearablesOnChainData, error) {
	var result wearables.WearablesOnChainData
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := c.db.WithContext(ctx).Where("id", id).Find(&result).Error; err != nil {
		c.l.Errorw("error get information of Wearables", "error", err, "wearables_id", id)
		return wearables.WearablesOnChainData{}, err
	}
	return result, nil
}

func (c *CockroachDB) getListWearablesOnchainData(offset int, limit int) ([]wearables.WearablesOnChainData, error) {
	return getListWearablesOnchainData(c.db, offset, limit)
}

func getListWearablesOnchainData(db *gorm.DB, offset int, limit int) ([]wearables.WearablesOnChainData, error) {
	var result []wearables.WearablesOnChainData
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := db.WithContext(ctx).Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return []wearables.WearablesOnChainData{}, err
	}
	return result, nil
}

func (c *CockroachDB) GetWearablesOffer(id uint64) (wearables.WearableOffers, error) {
	var result wearables.WearableOffers
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := c.db.WithContext(ctx).Where("on_chain_id", id).Find(&result).Error; err != nil {
		c.l.Errorw("error get information of Wearables Offer", "error", err, "onchain_id", id)
		return wearables.WearableOffers{}, err
	}
	return result, nil
}

func (c *CockroachDB) GetListWearablesOffer(offset int, limit int) ([]wearables.WearableOffers, uint64, error) {
	return getListWearablesOffer(c.db, offset, limit)
}

func getListWearablesOffer(db *gorm.DB, offset int, limit int) ([]wearables.WearableOffers, uint64, error) {
	var result []wearables.WearableOffers
	var count int64
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := db.WithContext(ctx).Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return []wearables.WearableOffers{}, 0, err
	}

	if err := db.WithContext(ctx).Table("wearable_offers").Count(&count).Error; err != nil {
		return []wearables.WearableOffers{}, 0, err
	}

	return result, uint64(count), nil
}

func (c *CockroachDB) GetWearablesListing(id uint64) (wearables.WearableListing, error) {
	var result wearables.WearableListing
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := c.db.WithContext(ctx).Where("on_chain_id", id).Find(&result).Error; err != nil {
		c.l.Errorw("error get information of Wearables Offer", "error", err, "onchain_id", id)
		return wearables.WearableListing{}, err
	}
	return result, nil
}

func (c *CockroachDB) GetListWearablesListing(offset int, limit int) ([]wearables.WearableListing, uint64, error) {
	return getListWearablesListing(c.db, offset, limit)
}

func getListWearablesListing(db *gorm.DB, offset int, limit int) ([]wearables.WearableListing, uint64, error) {
	var result []wearables.WearableListing
	var count int64
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := db.WithContext(ctx).Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return []wearables.WearableListing{}, 0, err
	}

	if err := db.WithContext(ctx).Table("wearable_listings").Count(&count).Error; err != nil {
		return []wearables.WearableListing{}, 0, err
	}
	return result, uint64(count), nil
}

func (c *CockroachDB) GetListWearablesActivity(offset int, limit int) ([]wearables.WearableActivities, uint64, error) {
	return getListWearablesActivity(c.db, offset, limit)
}

func getListWearablesActivity(db *gorm.DB, offset int, limit int) ([]wearables.WearableActivities, uint64, error) {
	var result []wearables.WearableActivities
	var count int64
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	if err := db.WithContext(ctx).Offset(offset).Limit(limit).Find(&result).Error; err != nil {
		return []wearables.WearableActivities{}, 0 , err
	}

	if err := db.WithContext(ctx).Table("wearable_activities").Count(&count).Error; err != nil {
		return []wearables.WearableActivities{}, 0 , err
	}

	return result, uint64(count), nil
}

func (c *CockroachDB) GetListWearablesInformation(offset int, limit int, filter wearables.WearablesFilterParams) ([]wearables.WearableInfo, uint64, error) {
	var listwearablesModel []wearables.WearableInfoModel
	var result []wearables.WearableInfo
	var count int64

	var sortedBy = "on_chain_id ASC"
	switch filter.Sort {
	case wearables.WearablesSortingLowestCurrentSupply:
		sortedBy = "current_supply ASC"
	case wearables.WearablesSortingHighestCurrentSupply:
		sortedBy = "current_supply DESC"
	case wearables.WearablesSortingLowestMaxSupply:
		sortedBy = "max_supply ASC"
	case wearables.WearablesSortingHighestMaxSupply:
		sortedBy = "max_supply DESC"
	}

	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	query := c.db.WithContext(ctx).Table("wearables").Select(
		"wearables.wearable_id, wearables.on_chain_id, wearables_on_chain_data.name, wearables.gender, wearables.changeable , wearables.is_origin , wearables.is_required ,wearables.image ,wearables.creativity , wearables.charisma , wearables.resolve , wearables.fitness , wearables.intellect , wearables_on_chain_data.max_supply ,wearables_on_chain_data.current_supply ,wearables.type ,wearables_on_chain_data.rarity ,wearables_on_chain_data.created_at_timestamp ,wearables_on_chain_data.updated_at_timestamp ,wearables_on_chain_data.created_at_block ,wearables_on_chain_data.updated_at_block").
		Joins("left join wearables_on_chain_data on wearables.on_chain_id = wearables_on_chain_data.id")
	if filter.Gender != nil && len(filter.Gender) > 0 {
		query.Where("wearables.gender IN (?)", filter.Gender)
	}

	if filter.Type != nil && len(filter.Type) > 0 {
		query.Where("wearables.type IN (?)", filter.Type)
	}

	if filter.Rarity != "" {
		query.Where("wearables_on_chain_data.rarity", filter.Rarity)
	}

	if err := query.Count(&count).Error; err != nil {
		c.l.Errorw("error get list wearables count", "error", err)
		return nil, 0, err
	}

	if err := query.Offset(offset * limit).Limit(limit).Order(sortedBy).Scan(&listwearablesModel).Error; err != nil {
		c.l.Errorw("error query get list wearable", "error", err)
		return nil, 0, err
	}

	for _,v := range listwearablesModel {
		stats := services.Stats {
			Creativity: v.Creativity,
			Charisma: v.Charisma,
			Resolve: v.Resolve,
			Fitness: v.Fitness,
			Intellect: v.Intellect,
		}
		wearable := wearables.WearableInfo {
			WearableId: v.WearableId,
			OnChainId: v.OnChainId,
			Name: v.Name,
			Gender: v.Gender,
			Changeable: v.Changeable,
			IsOrigin: v.IsOrigin,
			IsRequired: v.IsRequired,
			Image: v.Image,
			Stats: stats,
			MaxSupply: v.MaxSupply,
			CurrentSupply: v.CurrentSupply,
			Type: v.Type,
			Rarity: v.Rarity,
			CreatedAtTimestamp: v.CreatedAtTimestamp,
			UpdatedAtTimestamp: v.UpdatedAtTimestamp,
			CreatedAtBlock: v.CreatedAtBlock,
			UpdatedAtBlock: v.UpdatedAtBlock,
		}
		result = append(result, wearable)
	}
	return result, uint64(count), nil
}

func (c *CockroachDB) GetWearablesInformation(id uint64) (wearables.WearableInfo, error) {
	var v wearables.WearableInfoModel
	ctx, cancel := context.WithTimeout(context.Background(), services.DefaultCockroachDbTimeout)
	defer cancel()
	query := c.db.WithContext(ctx).Table("wearables").Select(
		"wearables.wearable_id, wearables.on_chain_id, wearables_on_chain_data.name, wearables.gender, wearables.changeable , wearables.is_origin , wearables.is_required ,wearables.image ,wearables.creativity , wearables.charisma , wearables.resolve , wearables.fitness , wearables.intellect , wearables_on_chain_data.max_supply ,wearables_on_chain_data.current_supply ,wearables.type ,wearables_on_chain_data.rarity ,wearables_on_chain_data.created_at_timestamp ,wearables_on_chain_data.updated_at_timestamp ,wearables_on_chain_data.created_at_block ,wearables_on_chain_data.updated_at_block").
		Joins("left join wearables_on_chain_data on wearables.on_chain_id = wearables_on_chain_data.id")
	if err := query.Where("on_chain_id", id).Scan(&v).Error; err != nil {
		return wearables.WearableInfo{}, err
	}

	stats := services.Stats {
		Creativity: v.Creativity,
		Charisma: v.Charisma,
		Resolve: v.Resolve,
		Fitness: v.Fitness,
		Intellect: v.Intellect,
	}
	result := wearables.WearableInfo {
		WearableId: v.WearableId,
		OnChainId: v.OnChainId,
		Name: v.Name,
		Gender: v.Gender,
		Changeable: v.Changeable,
		IsOrigin: v.IsOrigin,
		IsRequired: v.IsRequired,
		Image: v.Image,
		Stats: stats,
		MaxSupply: v.MaxSupply,
		CurrentSupply: v.CurrentSupply,
		Type: v.Type,
		Rarity: v.Rarity,
		CreatedAtTimestamp: v.CreatedAtTimestamp,
		UpdatedAtTimestamp: v.UpdatedAtTimestamp,
		CreatedAtBlock: v.CreatedAtBlock,
		UpdatedAtBlock: v.UpdatedAtBlock,
	}
	return result, nil
}
