package usecase

import (
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	eprops "github.com/nuttchai/go-ddd/internal/domain/entities/props"
	vprops "github.com/nuttchai/go-ddd/internal/domain/value-objects/props"
	httpdtos "github.com/nuttchai/go-ddd/internal/http/dtos"
)

func userFromCreateDTO(payload *httpdtos.CreateUserDTO) *entity.User {
	props := &eprops.UserProps{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Address: vprops.AddressProps{
			Street:  payload.Address.Street,
			City:    payload.Address.City,
			State:   payload.Address.State,
			ZipCode: payload.Address.ZipCode,
		},
	}
	return entity.NewUser(props)
}

func userFromUpdateDTO(payload *httpdtos.UpdateUserDTO) *entity.User {
	props := &eprops.UserProps{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Address: vprops.AddressProps{
			Street:  payload.Address.Street,
			City:    payload.Address.City,
			State:   payload.Address.State,
			ZipCode: payload.Address.ZipCode,
		},
	}
	return entity.NewUser(props, payload.ID)
}

func userToResponseDTO(u *entity.User) *httpdtos.UserDTO {
	return &httpdtos.UserDTO{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Address: httpdtos.AddressDTO{
			Street:  u.Address.Street,
			City:    u.Address.City,
			State:   u.Address.State,
			ZipCode: u.Address.ZipCode,
		},
	}
}
