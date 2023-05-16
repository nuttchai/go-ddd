package repository

import (
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
)

type IUserRepository interface {
	FindOneById(id string) (*entity.User, error)
	Save(entity *entity.User) error
	Delete(entity *entity.User) error
	IsExisted(entity *entity.User) bool
	UpdateFirstNameIfIdExist(id, firstName string) error
	FindOneByEmail(email string) (*entity.User, error)
}
