package application

import (
	http "github.com/nuttchai/go-ddd/common/http"
	cmapper "github.com/nuttchai/go-ddd/common/infrastructure/data-mappers"
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	service "github.com/nuttchai/go-ddd/internal/domain/services"
	dto "github.com/nuttchai/go-ddd/internal/dtos"
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

	userDTO := a.UserReqDataMapper.ToDalEntity(user)
	jsonOk := http.SuccessResponse(userDTO)
	return &http.APIResponse{APISuccess: jsonOk}
}