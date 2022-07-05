package cockroachdb

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewCockroachDbConnection() (*gorm.DB, error) {
	uri := viper.GetString("METROGALAXY_API_COCKROACH_DB_URI")
	logLevel := logger.Silent
	switch viper.GetString("METROGALAXY_API_COCKROACH_DB_LOG") {
	case "info":
		logLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}
