package app

import (
	api "github.com/nuttchai/go-ddd/common/api"
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

func (a *UserAppService) FindUserById(payload dto.FindUserByIdDTO) (*api.APISuccess, *api.APIError) {
	user, err := a.userService.FindOneById(payload.Id)
	if err != nil {
		jsonErr := api.BadRequestError(err)
		return nil, jsonErr
	}

	userDTO, err := a.userAppDataMapper.ToDalEntity(user)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return nil, jsonErr
	}

	jsonOk := api.SuccessResponse(userDTO)
	return jsonOk, nil
}
