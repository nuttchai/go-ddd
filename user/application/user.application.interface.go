package application

import (
	http "github.com/nuttchai/go-ddd/common/http"
	dto "github.com/nuttchai/go-ddd/user/dtos"
)

type IUserApplicationService interface {
	FindUserById(payload *dto.FindUserByIdDTO) *http.APIResponse
}
