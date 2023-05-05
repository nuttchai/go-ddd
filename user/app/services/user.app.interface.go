package app

import (
	http "github.com/nuttchai/go-ddd/common/http"
	dto "github.com/nuttchai/go-ddd/user/dtos"
)

type IUserAppService interface {
	FindUserById(payload dto.FindUserByIdDTO) (*http.APISuccess, *http.APIError)
}
