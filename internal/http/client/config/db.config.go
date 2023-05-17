package config

import (
	"errors"

	types "github.com/nuttchai/go-ddd/internal/shared/types"
	context "github.com/nuttchai/go-ddd/utils/context"
	env "github.com/nuttchai/go-ddd/utils/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	initConnectionTimeout   = 5
	dBTypeAssertionErrorMsg = "invalid_db_type_assertion"
)

func initDB() error {
	appEnv := types.AppConfig.GetENV()
	dbConfig := &gorm.Config{}
	if appEnv == env.Production.Name {
		dbConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	dsn := types.AppConfig.GetDBConfig().GetDSN()
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

	types.AppConfig.SetDBInstance(db)
	return nil
}

func getDB() (*gorm.DB, error) {
	dbInterface := types.AppConfig.GetDBConfig().GetDBInstance()
	db, ok := dbInterface.(*gorm.DB)
	if !ok {
		return nil, errors.New(dBTypeAssertionErrorMsg)
	}

	return db, nil
}
