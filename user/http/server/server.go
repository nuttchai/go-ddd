package http

import (
	appmapper "github.com/nuttchai/go-ddd/user/app/data-mappers"
	application "github.com/nuttchai/go-ddd/user/app/services"
	service "github.com/nuttchai/go-ddd/user/domain/services"
	controller "github.com/nuttchai/go-ddd/user/http/controllers"
	inframapper "github.com/nuttchai/go-ddd/user/infra/data-mappers"
	repository "github.com/nuttchai/go-ddd/user/infra/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitServer() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	userRepo := repository.NewUserRepository(db, &inframapper.UserDataMapper{})
	userRepo.FindOneByEmail("")
	userSvc := service.NewUserService(userRepo)
	userApp := application.NewUserApplicationService(userSvc, &appmapper.UserAppDataMapper{})
	_ = controller.NewUserController(userApp)
}
