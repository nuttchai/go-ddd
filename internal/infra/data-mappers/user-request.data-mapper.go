package data_mapper

import (
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	value_object "github.com/nuttchai/go-ddd/internal/domain/value-objects"
	dto "github.com/nuttchai/go-ddd/internal/dtos"
)

type UserRequestDataMapper struct{}

func (m *UserRequestDataMapper) ToDomainEntity(dalEntity *dto.UserDTO) *entity.User {
	props := &entity.UserProps{
		FirstName: dalEntity.FirstName,
		LastName:  dalEntity.LastName,
		Email:     dalEntity.Email,
		Address: value_object.Address{
			Street:  dalEntity.Address.Street,
			City:    dalEntity.Address.City,
			State:   dalEntity.Address.State,
			ZipCode: dalEntity.Address.ZipCode,
		},
	}
	return entity.NewUser(props, dalEntity.Id)

}

func (m *UserRequestDataMapper) ToDalEntity(domainEntity *entity.User) *dto.UserDTO {
	return &dto.UserDTO{
		Id:        domainEntity.Id,
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
