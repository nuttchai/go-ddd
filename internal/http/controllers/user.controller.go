package controller

import (
	"github.com/labstack/echo"
	api "github.com/nuttchai/go-ddd/common/http"
	http "github.com/nuttchai/go-ddd/common/http"
	application "github.com/nuttchai/go-ddd/internal/app"
	dto "github.com/nuttchai/go-ddd/internal/shared/dtos"
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
	payload := new(dto.FindUserByIdDTO)
	payload.ID = e.Param("id")
	if ok, err := payload.IsDTOValid(); !ok {
		jsonErr := http.BadRequestError(err)
		return e.JSON(jsonErr.Status, jsonErr)
	}

	result := c.UserApplicationService.FindUserById(payload)
	return e.JSON(result.Status(), result.Value())
}

func (c *UserController) CreateUser(e echo.Context) error {
	payload := new(dto.CreateUserDTO)
	if err := api.DecodeDTO(e, payload); err != nil {
		jsonErr := http.BadRequestError(err)
		return e.JSON(jsonErr.Status, jsonErr)
	}

	result := c.UserApplicationService.CreateUser(payload)
	return e.JSON(result.Status(), result.Value())
}

func (c *UserController) UpdateUser(e echo.Context) error {
	payload := new(dto.UpdateUserDTO)
	payload.ID = e.Param("id")
	if err := api.DecodeDTO(e, payload); err != nil {
		jsonErr := http.BadRequestError(err)
		return e.JSON(jsonErr.Status, jsonErr)
	}

	result := c.UserApplicationService.UpdateUser(payload)
	return e.JSON(result.Status(), result.Value())
}
