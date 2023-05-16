package router

import (
	"github.com/labstack/echo"
	url "github.com/nuttchai/go-ddd/internal/http/client/utils"
	controller "github.com/nuttchai/go-ddd/internal/http/controllers"
)

var builder *url.UrlBuilder

func init() {
	builder = url.NewUrlBuilder("user")
}

func InitUserRouter(e *echo.Echo, handler controller.IUserController) {
	e.GET(builder.BuildPath("/:id"), handler.FindUserById)
	e.POST(builder.BuildPath(), handler.CreateUser)
	e.PUT(builder.BuildPath("/:id"), handler.UpdateUser)
}
