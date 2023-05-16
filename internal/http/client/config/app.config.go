package config

import (
	"github.com/labstack/echo"
	application "github.com/nuttchai/go-ddd/internal/app"
	service "github.com/nuttchai/go-ddd/internal/domain/services"
	router "github.com/nuttchai/go-ddd/internal/http/client/routers"
	controller "github.com/nuttchai/go-ddd/internal/http/controllers"
	mapper "github.com/nuttchai/go-ddd/internal/infra/data-mappers"
	repository "github.com/nuttchai/go-ddd/internal/infra/repositories"
)

func initApp(e *echo.Echo) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	userMapper := mapper.NewUserDataMapper()
	userRequestMapper := mapper.NewUserRequestDataMapper()

	userRepo := repository.NewUserRepository(db, userMapper)
	userSvc := service.NewUserService(userRepo)
	userApp := application.NewUserApplicationService(userSvc, userRequestMapper)
	userHttp := controller.NewUserController(userApp)

	router.InitUserRouter(e, userHttp)

	return nil
}
