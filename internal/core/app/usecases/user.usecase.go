package usecase

import (
	sharederror "github.com/nuttchai/go-ddd/internal/common/errors"
	repository "github.com/nuttchai/go-ddd/internal/core/app/repositories"
	entity "github.com/nuttchai/go-ddd/internal/core/domain/entities"
	eprops "github.com/nuttchai/go-ddd/internal/core/domain/entities/props"
	vprops "github.com/nuttchai/go-ddd/internal/core/domain/value-objects/props"
	dto "github.com/nuttchai/go-ddd/internal/core/http/dtos"
)

type UserUsecase struct {
	userRepo repository.IUserRepository
}

func NewUserUsecase(userRepo repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (a *UserUsecase) FindUserById(payload *dto.FindUserByIdDTO) (*dto.UserDTO, error) {
	user, err := a.userRepo.FindOneById(payload.ID)
	if err != nil {
		return nil, err
	}

	userDto := &dto.UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Address: dto.AddressDTO{
			Street:  user.Address.Street,
			City:    user.Address.City,
			State:   user.Address.State,
			ZipCode: user.Address.ZipCode,
		},
	}
	return userDto, nil
}

func (a *UserUsecase) CreateUser(payload *dto.CreateUserDTO) error {
	user := entity.NewUser(&eprops.UserProps{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Address: vprops.AddressProps{
			Street:  payload.Address.Street,
			City:    payload.Address.City,
			State:   payload.Address.State,
			ZipCode: payload.Address.ZipCode,
		},
	})
	if isUserValid := user.IsUserValid(); !isUserValid {
		return sharederror.ErrInvalidCreatedUser
	}

	existing, err := a.userRepo.FindOneByEmail(user.Email)
	if err != nil {
		return err
	}
	if existing != nil {
		return sharederror.ErrEmailAlreadyExisted
	}

	return a.userRepo.Save(user)
}

func (a *UserUsecase) UpdateUser(payload *dto.UpdateUserDTO) error {
	user := entity.NewUser(&eprops.UserProps{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Address: vprops.AddressProps{
			Street:  payload.Address.Street,
			City:    payload.Address.City,
			State:   payload.Address.State,
			ZipCode: payload.Address.ZipCode,
		},
	}, payload.ID)
	if isUserValid := user.IsUserValid(); !isUserValid {
		return sharederror.ErrInvalidUpdatedUser
	}

	return a.userRepo.UpdateUser(user)
}
