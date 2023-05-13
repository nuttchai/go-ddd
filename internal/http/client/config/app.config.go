package config

import (
	"github.com/labstack/echo"
	application "github.com/nuttchai/go-ddd/internal/app"
	service "github.com/nuttchai/go-ddd/internal/domain/services"
	controller "github.com/nuttchai/go-ddd/internal/http/controllers"
	router "github.com/nuttchai/go-ddd/internal/http/routers"
	mapper "github.com/nuttchai/go-ddd/internal/infra/data-mappers"
	repository "github.com/nuttchai/go-ddd/internal/infra/repositories"
)

func initApp(e *echo.Echo) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	userRepo := repository.NewUserRepository(db, &mapper.UserDataMapper{})
	userSvc := service.NewUserService(userRepo)
	userApp := application.NewUserApplicationService(userSvc, &mapper.UserRequestDataMapper{})
	userHttp := controller.NewUserController(userApp)

	router.InitUserRouter(e, userHttp)
	return nil
}
