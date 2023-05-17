package router

import (
	"github.com/labstack/echo"
	controller "github.com/nuttchai/go-ddd/internal/http/controllers"
)

func (r *Router) InitUserRouter(e *echo.Echo, handler controller.IUserController) {
	e.GET(r.userUrlBuilder.buildPath("/:id"), handler.FindUserById)
	e.PUT(r.userUrlBuilder.buildPath("/:id"), handler.UpdateUser)
	e.POST(r.userUrlBuilder.buildPath(), handler.CreateUser)
}
