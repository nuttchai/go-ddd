package dto

import validator "github.com/nuttchai/go-ddd/utils/validators"

type UpdateUserDTO struct {
	ID        string     `param:"id" validate:"required"`
	FirstName string     `json:"first_name" validate:"required"`
	LastName  string     `json:"last_name" validate:"required"`
	Email     string     `json:"email" validate:"required"`
	Address   AddressDTO `json:"address"`
}

func (dto *UpdateUserDTO) IsDTOValid() (bool, error) {
	return validator.IsValidStruct(dto)
}
