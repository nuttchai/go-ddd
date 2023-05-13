package router

import (
	"github.com/labstack/echo"
	controller "github.com/nuttchai/go-ddd/internal/presentation/controllers"
)

func InitUserRouter(e *echo.Echo, handler controller.IUserController) {
	e.GET(buildPath("/internal/:id"), handler.FindUserById)
}
