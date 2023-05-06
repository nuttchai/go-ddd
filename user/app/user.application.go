package app

import (
	api "github.com/nuttchai/go-ddd/common/api"
	cmapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	entity "github.com/nuttchai/go-ddd/user/domain/entities"
	service "github.com/nuttchai/go-ddd/user/domain/services"
	dto "github.com/nuttchai/go-ddd/user/dtos"
)

type UserApplicationService struct {
	userService       service.IUserService
	UserReqDataMapper cmapper.IDataMapper[entity.User, dto.UserDTO]
}

func NewUserApplicationService(userService service.IUserService, UserReqDataMapper cmapper.IDataMapper[entity.User, dto.UserDTO]) IUserApplicationService {
	return &UserApplicationService{
		userService:       userService,
		UserReqDataMapper: UserReqDataMapper,
	}
}

func (a *UserApplicationService) FindUserById(payload *dto.FindUserByIdDTO) *api.APIResponse {
	user, err := a.userService.FindOneById(payload.Id)
	if err != nil {
		jsonErr := api.BadRequestError(err)
		return &api.APIResponse{APIError: jsonErr}
	}

	userDTO, err := a.UserReqDataMapper.ToDalEntity(user)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return &api.APIResponse{APIError: jsonErr}
	}

	jsonOk := api.SuccessResponse(userDTO)
	return &api.APIResponse{APISuccess: jsonOk}
}
