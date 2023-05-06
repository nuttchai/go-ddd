package config

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-ddd/user/application"
	service "github.com/nuttchai/go-ddd/user/domain/services"
	mapper "github.com/nuttchai/go-ddd/user/infrastructure/data-mappers"
	repository "github.com/nuttchai/go-ddd/user/infrastructure/repositories"
	controller "github.com/nuttchai/go-ddd/user/presentation/controllers"
	router "github.com/nuttchai/go-ddd/user/presentation/routers"
)

func initApp(e *echo.Echo) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	userRepo := repository.NewUserRepository(db, &mapper.UserDataMapper{})
	userSvc := service.NewUserService(userRepo)
	userApp := application.NewUserApplicationService(userSvc, &mapper.UserReqDataMapper{})
	userHttp := controller.NewUserController(userApp)

	router.InitUserRouter(e, userHttp)
	return nil
}
