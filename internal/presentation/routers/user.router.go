package router

import (
	path "github.com/nuttchai/go-ddd/internal/presentation/routers/path"

	"github.com/labstack/echo"
	controller "github.com/nuttchai/go-ddd/internal/presentation/controllers"
)

func InitUserRouter(e *echo.Echo, handler controller.IUserController) {
	e.GET(path.CreatePath("/internal/:id"), handler.FindUserById)
}
