package usecase

import dto "github.com/nuttchai/go-ddd/internal/core/http/dtos"

type IUserUsecase interface {
	FindUserById(payload *dto.FindUserByIdDTO) (*dto.UserDTO, error)
	CreateUser(payload *dto.CreateUserDTO) error
	UpdateUser(payload *dto.UpdateUserDTO) error
}
