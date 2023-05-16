package controller

import "github.com/labstack/echo"

type IUserController interface {
	FindUserById(e echo.Context) error
	CreateUser(e echo.Context) error
	UpdateUser(e echo.Context) error
}
