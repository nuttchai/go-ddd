package data_mapper

import (
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	"github.com/nuttchai/go-ddd/internal/domain/props"
	model "github.com/nuttchai/go-ddd/internal/infra/models"
)

type UserDataMapper struct{}

func (m *UserDataMapper) ToDomainEntity(dalEntity *model.User) *entity.User {
	props := &props.UserProps{
		FirstName: dalEntity.FirstName,
		LastName:  dalEntity.LastName,
		Email:     dalEntity.Email,
		Address: props.AddressProps{
			Street:  dalEntity.Address.Street,
			City:    dalEntity.Address.City,
			State:   dalEntity.Address.State,
			ZipCode: dalEntity.Address.ZipCode,
		},
	}
	return entity.NewUser(props, dalEntity.Id)

}

func (m *UserDataMapper) ToDalEntity(domainEntity *entity.User) *model.User {
	return &model.User{
		Id:        domainEntity.Id,
		FirstName: domainEntity.FirstName,
		LastName:  domainEntity.LastName,
		Email:     domainEntity.Email,
		Address: model.Address{
			Street:  domainEntity.Address.Street,
			City:    domainEntity.Address.City,
			State:   domainEntity.Address.State,
			ZipCode: domainEntity.Address.ZipCode,
		},
	}
}
