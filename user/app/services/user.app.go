package app

import (
	http "github.com/nuttchai/go-ddd/common/http"
	cmapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	entity "github.com/nuttchai/go-ddd/user/domain/entities"
	service "github.com/nuttchai/go-ddd/user/domain/services"
	dto "github.com/nuttchai/go-ddd/user/dtos"
)

type UserAppService struct {
	userService       service.IUserService
	userAppDataMapper cmapper.IDataMapper[entity.User, dto.UserDTO]
}

func NewUserAppService(userService service.IUserService, userAppDataMapper cmapper.IDataMapper[entity.User, dto.UserDTO]) *UserAppService {
	return &UserAppService{
		userService:       userService,
		userAppDataMapper: userAppDataMapper,
	}
}

func (a *UserAppService) FindUserById(payload dto.FindUserByIdDTO) (*http.APISuccess, *http.APIError) {
	user, err := a.userService.FindOneById(payload.Id)
	if err != nil {
		jsonErr := http.BadRequestError(err)
		return nil, jsonErr
	}

	userDTO, err := a.userAppDataMapper.ToDalEntity(user)
	if err != nil {
		jsonErr := http.InternalServerError(err)
		return nil, jsonErr
	}

	jsonOk := http.SuccessResponse(userDTO)
	return jsonOk, nil
}
