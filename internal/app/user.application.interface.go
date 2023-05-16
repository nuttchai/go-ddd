package application

import (
	http "github.com/nuttchai/go-ddd/common/http"
	dto "github.com/nuttchai/go-ddd/internal/shared/dtos"
)

type IUserApplicationService interface {
	FindUserById(payload *dto.FindUserByIdDTO) *http.APIResponse
	CreateUser(payload *dto.CreateUserDTO) *http.APIResponse
	UpdateUser(payload *dto.UpdateUserDTO) *http.APIResponse
}
