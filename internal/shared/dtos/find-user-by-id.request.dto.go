package dto

import (
	"errors"

	validator "github.com/nuttchai/go-ddd/utils/validators"
)

type FindUserByIdDTO struct {
	Id string `query:"id" validate:"required"`
}

func (dto *FindUserByIdDTO) IsDTOValid() (bool, error) {
	if isUUID := validator.IsUUID(dto.Id); !isUUID {
		return false, errors.New("invalid_uuid")
	}
	return validator.IsValidStruct(dto)
}
