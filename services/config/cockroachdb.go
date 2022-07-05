package config

import (
	"context"
	"fmt"
	"time"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	DefaultTimeout  = 30 * time.Second
	DefaultConfigID = 1
)

type CockroachDB struct {
	l  *zap.SugaredLogger
	db *gorm.DB
}

func NewCockroachDB(db *gorm.DB) (*CockroachDB, error) {
	tables := make([]interface{}, 0)
	tables = append(tables, &Config{})

	if err := db.AutoMigrate(tables...); err != nil {
		return nil, err
	}

	return &CockroachDB{
		l:  zap.S(),
		db: db,
	}, nil
}

func (c *CockroachDB) SaveConfig(config Config) error {
	config.ID = DefaultConfigID
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return tx.Clauses(clause.OnConflict{
				UpdateAll: true,
			}).Create(&config).Error
		},
	); err != nil {
		c.l.Errorw("error save config", "error", err)
		return err
	}
	return nil
}

// UpdateConfig upsert
func (c *CockroachDB) UpdateConfig(params map[ConfigField]interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	if err := crdbgorm.ExecuteTx(ctx, c.db, nil,
		func(tx *gorm.DB) error {
			return updateConfig(tx, params)
		},
	); err != nil {
		c.l.Errorw("error update config", "error", err, "params", params)
		return err
	}
	return nil
}

func updateConfig(db *gorm.DB, params map[ConfigField]interface{}) error {
	data := &Config{
		ID: DefaultConfigID,
	}
	fieldsToUpdate := make([]string, 0)

	for key, value := range params {
		if key == ConfigFieldLastFetchedBlock {
			var ok bool
			data.LastFetchedBlock, ok = value.(uint64)
			if !ok {
				return fmt.Errorf("error convert field `last_fetched_block`")
			}
			fieldsToUpdate = append(fieldsToUpdate, string(key))
		}
	}

	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(fieldsToUpdate),
	}).Create(&data).Error
}

func (c *CockroachDB) GetConfig() (Config, error) {
	var result Config
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()
	if err := c.db.WithContext(ctx).Where("id = ?", DefaultConfigID).Take(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return Config{}, nil
		}
		c.l.Errorw("error get config", "error", err)
		return Config{}, err
	}

	return result, nil
}
