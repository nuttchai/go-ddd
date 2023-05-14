package router

import (
	"github.com/labstack/echo"
	controller "github.com/nuttchai/go-ddd/internal/http/controllers"
)

func InitUserRouter(e *echo.Echo, handler controller.IUserController) {
	e.GET(buildPath("user/:id"), handler.FindUserById)
	e.POST(buildPath("user"), handler.CreateUser)
}
