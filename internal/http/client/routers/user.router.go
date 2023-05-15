package router

import (
	"github.com/labstack/echo"
	url "github.com/nuttchai/go-ddd/internal/http/client/utils"
	controller "github.com/nuttchai/go-ddd/internal/http/controllers"
)

func InitUserRouter(e *echo.Echo, handler controller.IUserController) {
	e.GET(url.BuildPath("user/:id"), handler.FindUserById)
	e.POST(url.BuildPath("user"), handler.CreateUser)
}
