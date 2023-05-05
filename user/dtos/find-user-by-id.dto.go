package dto

import (
	"errors"

	constant "github.com/nuttchai/go-ddd/common/constants"
	uuid "github.com/nuttchai/go-ddd/utils/validators"
)

type FindUserByIdDTO struct {
	Id string `json:"id"`
}

func NewFindUserByIdDTO(id string) (*FindUserByIdDTO, error) {
	if isUUID := uuid.IsUUID(id); !isUUID {
		return nil, errors.New(constant.InvalidUUID)
	}

	return &FindUserByIdDTO{
		Id: id,
	}, nil
}
