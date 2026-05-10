package config

import (
	"github.com/labstack/echo"
	usecase "github.com/nuttchai/go-ddd/internal/app/usecases"
	route "github.com/nuttchai/go-ddd/internal/http/client/routers"
	controller "github.com/nuttchai/go-ddd/internal/http/controllers"
	mapper "github.com/nuttchai/go-ddd/internal/infra/data-mappers"
	repository "github.com/nuttchai/go-ddd/internal/infra/repositories"
)

func initApp(e *echo.Echo) error {
	router := route.NewRouter()
	if err := initUserApp(e, router); err != nil {
		return err
	}

	return nil
}

func initUserApp(e *echo.Echo, r *route.Router) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	userMapper := mapper.NewUserDataMapper()

	repo := repository.NewUserRepository(db, userMapper)
	userUsecase := usecase.NewUserUsecase(repo)
	http := controller.NewUserController(userUsecase)

	r.InitUserRouter(e, http)

	return nil
}
