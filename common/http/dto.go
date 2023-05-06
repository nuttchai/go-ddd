package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/labstack/echo"
	validator "github.com/nuttchai/go-ddd/utils/validators"
)

func DecodeDTO(e echo.Context, ptr any) error {
	decoder := json.NewDecoder(e.Request().Body)
	if err := decoder.Decode(ptr); err != nil {
		msg := fmt.Sprintf("decoding json error: %s", err.Error())
		return errors.New(msg)
	}

	if _, err := validator.IsValidStruct(ptr); err != nil {
		msg := fmt.Sprintf("validating dto error: %s", err.Error())
		return errors.New(msg)
	}

	return nil
}
