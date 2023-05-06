package data_mapper

import (
	entity "github.com/nuttchai/go-ddd/user/domain/entities"
	value_object "github.com/nuttchai/go-ddd/user/domain/value-objects"
	dto "github.com/nuttchai/go-ddd/user/dtos"
)

type UserReqDataMapper struct{}

func (m *UserReqDataMapper) ToDomainEntity(dalEntity *dto.UserDTO) (*entity.User, error) {
	return &entity.User{
		Id:        dalEntity.Id,
		FirstName: dalEntity.FirstName,
		LastName:  dalEntity.LastName,
		Email:     dalEntity.Email,
		Address: value_object.Address{
			Street:  dalEntity.Address.Street,
			City:    dalEntity.Address.City,
			State:   dalEntity.Address.State,
			ZipCode: dalEntity.Address.ZipCode,
		},
	}, nil
}

func (m *UserReqDataMapper) ToDalEntity(domainEntity *entity.User) (*dto.UserDTO, error) {
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
	}, nil
}
