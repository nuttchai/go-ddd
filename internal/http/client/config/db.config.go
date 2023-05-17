package config

import (
	"errors"

	context "github.com/nuttchai/go-ddd/utils/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	initConnectionTimeout   = 5
	dBTypeAssertionErrorMsg = "invalid_db_type_assertion"
)

func initDB() error {
	env := AppConfig.GetENV()
	dbConfig := &gorm.Config{}
	if env == "production" {
		dbConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	dsn := AppConfig.GetDBConfig().GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), dbConfig)
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(initConnectionTimeout)
	defer cancel()
	if err = sqlDB.PingContext(ctx); err != nil {
		return err
	}

	AppConfig.SetDBInstance(db)
	return nil
}

func getDB() (*gorm.DB, error) {
	dbInterface := AppConfig.GetDBConfig().GetDBInstance()
	db, ok := dbInterface.(*gorm.DB)
	if !ok {
		return nil, errors.New(dBTypeAssertionErrorMsg)
	}

	return db, nil
}
