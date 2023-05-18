package repository

import (
	repository "github.com/nuttchai/go-ddd/common/infra/repositories"
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	model "github.com/nuttchai/go-ddd/internal/infra/models"
)

type IUserRepository interface {
	repository.IRepository[entity.User, model.User]
	UpdateUser(entity *entity.User) error
	FindOneByEmail(email string) (*entity.User, error)
}
