package inject

import (
	cockroachdblib "github.com/metrogalaxy-io/metrogalaxy-api/libs/cockroachdb"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/config"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/metronion"
	metronionCockroachDb "github.com/metrogalaxy-io/metrogalaxy-api/services/metronion/cockroachdb"
	"github.com/metrogalaxy-io/metrogalaxy-api/services/wearables"
	wearableCockroachDb "github.com/metrogalaxy-io/metrogalaxy-api/services/wearables/cockroachdb"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (c *Injector) ProvideGormConnection() *gorm.DB {
	if c.gormConnection == nil {
		l := zap.S()
		connection, err := cockroachdblib.NewCockroachDbConnection()
		if err != nil {
			l.Panic(err)
		}
		c.gormConnection = connection
	}
	return c.gormConnection
}

func (c *Injector) ProvideMetronionDb() metronion.PersistentDb {
	if c.metronionDb == nil {
		l := zap.S()
		db, err := metronionCockroachDb.NewCockroachDB(c.ProvideGormConnection())
		if err != nil {
			l.Panic(err)
		}
		c.metronionDb = db
	}
	return c.metronionDb
}

func (c *Injector) ProvideConfigDb() config.PersistentDb {
	if c.configDb == nil {
		l := zap.S()
		db, err := config.NewCockroachDB(c.ProvideGormConnection())
		if err != nil {
			l.Panic(err)
		}
		c.configDb = db
	}
	return c.configDb
}

func (c *Injector) ProvideWearableDb() wearables.PersistentDb {
	if c.wearablesDb == nil {
		l := zap.S()
		db, err := wearableCockroachDb.NewCockroachDB(c.ProvideGormConnection())
		if err != nil {
			l.Panic(err)
		}
		c.wearablesDb = db
	}
	return c.wearablesDb
}
