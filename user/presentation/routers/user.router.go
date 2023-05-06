package router

import (
	path "github.com/nuttchai/go-ddd/user/presentation/routers/path"

	"github.com/labstack/echo"
	controller "github.com/nuttchai/go-ddd/user/presentation/controllers"
)

func InitUserRouter(e *echo.Echo, handler controller.IUserController) {
	e.GET(path.CreatePath("/user/:id"), handler.FindUserById)
}
