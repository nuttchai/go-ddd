package api

import (
	"github.com/labstack/echo"
)

type DTO interface {
	IsDTOValid() (bool, error)
}

func DecodeDTO(e echo.Context, dto DTO) error {
	if err := e.Bind(&dto); err != nil {
		return err
	}
	if ok, err := dto.IsDTOValid(); !ok {
		return err
	}
	return nil
}
