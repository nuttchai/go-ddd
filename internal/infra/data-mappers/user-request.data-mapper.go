package data_mapper

import (
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	eprops "github.com/nuttchai/go-ddd/internal/domain/entities/props"
	vprops "github.com/nuttchai/go-ddd/internal/domain/value-objects/props"
	dto "github.com/nuttchai/go-ddd/internal/shared/dtos"
)

type UserRequestDataMapper struct{}

func (m *UserRequestDataMapper) ToDomainEntity(dalEntity *dto.UserDTO) *entity.User {
	props := &eprops.UserProps{
		FirstName: dalEntity.FirstName,
		LastName:  dalEntity.LastName,
		Email:     dalEntity.Email,
		Address: vprops.AddressProps{
			Street:  dalEntity.Address.Street,
			City:    dalEntity.Address.City,
			State:   dalEntity.Address.State,
			ZipCode: dalEntity.Address.ZipCode,
		},
	}
	return entity.NewUser(props, dalEntity.ID)
}

func (m *UserRequestDataMapper) ToDalEntity(domainEntity *entity.User) *dto.UserDTO {
	return &dto.UserDTO{
		ID:        domainEntity.ID,
		FirstName: domainEntity.FirstName,
		LastName:  domainEntity.LastName,
		Email:     domainEntity.Email,
		Address: dto.AddressDTO{
			Street:  domainEntity.Address.Street,
			City:    domainEntity.Address.City,
			State:   domainEntity.Address.State,
			ZipCode: domainEntity.Address.ZipCode,
		},
	}
}

func NewUserRequestDataMapper() *UserRequestDataMapper {
	return &UserRequestDataMapper{}
}
