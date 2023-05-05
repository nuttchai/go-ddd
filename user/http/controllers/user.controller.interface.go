package controller

import "github.com/labstack/echo"

type IUserController interface {
	FindUserById(e echo.Context) error
}
