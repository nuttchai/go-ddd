package controller

import (
	"net/http"

	"github.com/labstack/echo"
	chttp "github.com/nuttchai/go-ddd/common/http"
	usecase "github.com/nuttchai/go-ddd/internal/app/usecases"
	dto "github.com/nuttchai/go-ddd/internal/http/dtos"
	"github.com/nuttchai/go-ddd/internal/shared/helpers"
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
		res := chttp.Err(err.Error(), http.StatusBadRequest, chttp.CodeBadRequestError)
		return e.JSON(res.Status, res.Envelope)
	}

	data, err := c.userUsecase.FindUserById(payload)
	res := helpers.BuildEnvelope(data, err)
	return e.JSON(res.Status, res.Envelope)
}

func (c *UserController) CreateUser(e echo.Context) error {
	payload := new(dto.CreateUserDTO)
	if err := chttp.DecodeDTO(e, payload); err != nil {
		res := chttp.Err(err.Error(), http.StatusBadRequest, chttp.CodeBadRequestError)
		return e.JSON(res.Status, res.Envelope)
	}

	err := c.userUsecase.CreateUser(payload)
	res := helpers.BuildEnvelope(nil, err)
	return e.JSON(res.Status, res.Envelope)
}

func (c *UserController) UpdateUser(e echo.Context) error {
	payload := new(dto.UpdateUserDTO)
	payload.ID = e.Param("id")
	if err := chttp.DecodeDTO(e, payload); err != nil {
		res := chttp.Err(err.Error(), http.StatusBadRequest, chttp.CodeBadRequestError)
		return e.JSON(res.Status, res.Envelope)
	}

	err := c.userUsecase.UpdateUser(payload)
	res := helpers.BuildEnvelope(nil, err)
	return e.JSON(res.Status, res.Envelope)
}
