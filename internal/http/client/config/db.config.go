package config

import (
	"errors"

	constant "github.com/nuttchai/go-ddd/internal/http/client/constants"
	context "github.com/nuttchai/go-ddd/utils/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() error {
	dsn := AppConfig.GetDBConfig().GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(constant.InitConnectionTimeout)
	defer cancel()
	if err = sqlDB.PingContext(ctx); err != nil {
		return err
	}

	defer sqlDB.Close()
	AppConfig.SetDBInstance(db)
	return nil
}

func getDB() (*gorm.DB, error) {
	dbInterface := AppConfig.GetDBConfig().GetDB()
	db, ok := dbInterface.(*gorm.DB)
	if !ok {
		return nil, errors.New(constant.InvalidDBTypeAssertion)
	}

	return db, nil
}
