package dto

import validator "github.com/nuttchai/go-ddd/utils/validators"

type CreateAddressDTO struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zipCode"`
}

type CreateUserDTO struct {
	FirstName string           `json:"first_name" validate:"required"`
	LastName  string           `json:"last_name" validate:"required"`
	Email     string           `json:"email" validate:"required"`
	Address   CreateAddressDTO `json:"address"`
}

func (dto *CreateUserDTO) IsDTOValid() (bool, error) {
	return validator.IsValidStruct(dto)
}
