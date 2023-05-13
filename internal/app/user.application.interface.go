package application

import (
	http "github.com/nuttchai/go-ddd/common/http"
	dto "github.com/nuttchai/go-ddd/internal/dtos"
)

type IUserApplicationService interface {
	FindUserById(payload *dto.FindUserByIdDTO) *http.APIResponse
}
