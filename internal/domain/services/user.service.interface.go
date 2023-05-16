package service

import entity "github.com/nuttchai/go-ddd/internal/domain/entities"

type IUserService interface {
	FindOneById(id string) (*entity.User, error)
	CreateUser(user *entity.User) error
	UpdateUser(user *entity.User) error
	FindOneByEmail(email string) (*entity.User, error)
}
