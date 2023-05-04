package service

import (
	"errors"

	constant "github.com/nuttchai/go-ddd/user/domain/constants"
	entity "github.com/nuttchai/go-ddd/user/domain/entities"
	repo "github.com/nuttchai/go-ddd/user/domain/repositories"
)

type UserService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) FindOneById(id string) (*entity.User, error) {
	return s.userRepo.FindOneById(id)
}

func (s *UserService) CreateUser(user *entity.User) error {
	if isUserValid := user.IsUserValid(); !isUserValid {
		return errors.New(constant.InvalidCreatedUser)
	}
	return s.userRepo.Save(user)
}

func (s *UserService) UpdateUser(user *entity.User) error {
	if isUserValid := user.IsUserValid(); !isUserValid {
		return errors.New(constant.InvalidUpdatedUser)
	}
	return s.userRepo.Update(user)
}

func (s *UserService) UpdateFirstNameIfIdExist(id, firstName string) error {
	return s.userRepo.UpdateFirstNameIfIdExist(id, firstName)
}

func (s *UserService) FindOneByEmail(email string) (*entity.User, error) {
	return s.userRepo.FindOneByEmail(email)
}
