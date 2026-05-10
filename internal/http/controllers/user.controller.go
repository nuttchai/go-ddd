package controller

import (
	"github.com/labstack/echo"
	http "github.com/nuttchai/go-ddd/common/http"
	usecase "github.com/nuttchai/go-ddd/internal/app/usecases"
	dto "github.com/nuttchai/go-ddd/internal/http/dtos"
)

type UserController struct {
	userUsecase usecase.IUserUsecase
}

func NewUserController(userUsecase usecase.IUserUsecase) IUserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (c *UserController) FindUserById(e echo.Context) error {
	payload := new(dto.FindUserByIdDTO)
	payload.ID = e.Param("id")
	if ok, err := payload.IsDTOValid(); !ok {
		jsonErr := http.BadRequestError(err)
		return e.JSON(jsonErr.Status, jsonErr)
	}

	result := c.userUsecase.FindUserById(payload)
	return e.JSON(result.Status(), result.Value())
}

func (c *UserController) CreateUser(e echo.Context) error {
	payload := new(dto.CreateUserDTO)
	if err := http.DecodeDTO(e, payload); err != nil {
		jsonErr := http.BadRequestError(err)
		return e.JSON(jsonErr.Status, jsonErr)
	}

	result := c.userUsecase.CreateUser(payload)
	return e.JSON(result.Status(), result.Value())
}

func (c *UserController) UpdateUser(e echo.Context) error {
	payload := new(dto.UpdateUserDTO)
	payload.ID = e.Param("id")
	if err := http.DecodeDTO(e, payload); err != nil {
		jsonErr := http.BadRequestError(err)
		return e.JSON(jsonErr.Status, jsonErr)
	}

	result := c.userUsecase.UpdateUser(payload)
	return e.JSON(result.Status(), result.Value())
}
