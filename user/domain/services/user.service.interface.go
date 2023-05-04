package service

import entity "github.com/nuttchai/go-ddd/user/domain/entities"

type IUserService interface {
	FindOneById(id string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	UpdateFirstNameIfIdExist(id, firstName string) error
	FindOneByEmail(email string) (*entity.User, error)
}
