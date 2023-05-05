package controller

import (
	"github.com/labstack/echo"
	"github.com/nuttchai/go-ddd/common/api"
	application "github.com/nuttchai/go-ddd/user/app/services"
	dto "github.com/nuttchai/go-ddd/user/dtos"
)

type UserController struct {
	UserApplicationService application.IUserApplicationService
}

func NewUserController(UserApplicationService application.IUserApplicationService) IUserController {
	return &UserController{
		UserApplicationService: UserApplicationService,
	}
}

func (c *UserController) FindUserById(e echo.Context) error {
	id := e.Param("id")
	payload, err := dto.NewFindUserByIdDTO(id)
	if err != nil {
		jsonErr := api.BadRequestError(err)
		return e.JSON(jsonErr.Status, jsonErr)
	}

	result := c.UserApplicationService.FindUserById(payload)
	return e.JSON(result.Status(), result.Value())
}
