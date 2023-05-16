package service

import (
	"errors"

	constant "github.com/nuttchai/go-ddd/internal/domain/constants"
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	repo "github.com/nuttchai/go-ddd/internal/domain/repositories"
)

type UserService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
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
	if userDb, err := s.userRepo.FindOneByEmail(user.Email); userDb != nil || err != nil {
		if err != nil {
			return err
		}
		return errors.New(constant.EmailAlreadyExisted)
	}
	return s.userRepo.Save(user)
}

func (s *UserService) UpdateUser(user *entity.User) error {
	if isUserValid := user.IsUserValid(); !isUserValid {
		return errors.New(constant.InvalidUpdatedUser)
	}
	return s.userRepo.UpdateUser(user)
}

func (s *UserService) FindOneByEmail(email string) (*entity.User, error) {
	return s.userRepo.FindOneByEmail(email)
}
