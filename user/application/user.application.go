package application

import (
	http "github.com/nuttchai/go-ddd/common/http"
	cmapper "github.com/nuttchai/go-ddd/common/infrastructure/data-mappers"
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

func (a *UserApplicationService) FindUserById(payload *dto.FindUserByIdDTO) *http.APIResponse {
	user, err := a.userService.FindOneById(payload.Id)
	if err != nil {
		jsonErr := http.BadRequestError(err)
		return &http.APIResponse{APIError: jsonErr}
	}

	userDTO, err := a.UserReqDataMapper.ToDalEntity(user)
	if err != nil {
		jsonErr := http.InternalServerError(err)
		return &http.APIResponse{APIError: jsonErr}
	}

	jsonOk := http.SuccessResponse(userDTO)
	return &http.APIResponse{APISuccess: jsonOk}
}
