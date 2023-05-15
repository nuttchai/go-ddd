package dto

import validator "github.com/nuttchai/go-ddd/utils/validators"

type CreateUserDTO struct {
	FirstName string     `json:"first_name" validate:"required"`
	LastName  string     `json:"last_name" validate:"required"`
	Email     string     `json:"email" validate:"required"`
	Address   AddressDTO `json:"address"`
}

func (dto *CreateUserDTO) IsDTOValid() (bool, error) {
	return validator.IsValidStruct(dto)
}
