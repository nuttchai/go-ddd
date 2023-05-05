package app

import (
	api "github.com/nuttchai/go-ddd/common/api"
	dto "github.com/nuttchai/go-ddd/user/dtos"
)

type IUserAppService interface {
	FindUserById(payload dto.FindUserByIdDTO) (*api.APISuccess, *api.APIError)
}
