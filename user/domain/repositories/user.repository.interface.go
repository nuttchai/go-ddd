package repository

import (
	entity "github.com/nuttchai/go-ddd/user/domain/entities"
)

type IUserRepository interface {
	FindOneById(id string) (*entity.User, error)
	Save(entity *entity.User) error
	Update(entity *entity.User) error
	Delete(id string) error
	IsExisted(entity *entity.User) (bool, error)
	UpdateFirstNameIfIdExist(id, firstName string) error
	FindOneByEmail(email string) (*entity.User, error)
}
