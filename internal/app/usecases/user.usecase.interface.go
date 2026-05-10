package usecase

import (
	http "github.com/nuttchai/go-ddd/common/http"
	httpdtos "github.com/nuttchai/go-ddd/internal/http/dtos"
)

type IUserUsecase interface {
	FindUserById(payload *httpdtos.FindUserByIdDTO) *http.APIResponse
	CreateUser(payload *httpdtos.CreateUserDTO) *http.APIResponse
	UpdateUser(payload *httpdtos.UpdateUserDTO) *http.APIResponse
}
