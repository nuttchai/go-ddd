package config

import (
	constant "github.com/nuttchai/go-ddd/user/config/constants"
	context "github.com/nuttchai/go-ddd/utils/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() (*gorm.DB, error) {
	dsn := AppConfig.GetDBConfig().GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(constant.InitConnectionTimeout)
	defer cancel()
	if err = sqlDB.PingContext(ctx); err != nil {
		return nil, err
	}

	defer sqlDB.Close()
	return db, nil
}
