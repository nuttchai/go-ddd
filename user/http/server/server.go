package http

import (
	"github.com/nuttchai/go-ddd/user/app"
	service "github.com/nuttchai/go-ddd/user/domain/services"
	mapper "github.com/nuttchai/go-ddd/user/infra/data-mappers"
	repository "github.com/nuttchai/go-ddd/user/infra/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitServer() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	userRepo := repository.NewUserRepository(db, &mapper.UserDataMapper{})
	userSvc := service.NewUserService(userRepo)
	_ = app.NewUserApp(userSvc)
}
