package application

import (
	http "github.com/nuttchai/go-ddd/common/http"
	cmapper "github.com/nuttchai/go-ddd/common/infra/data-mappers"
	entity "github.com/nuttchai/go-ddd/internal/domain/entities"
	service "github.com/nuttchai/go-ddd/internal/domain/services"
	dto "github.com/nuttchai/go-ddd/internal/shared/dtos"
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
	user, err := a.userService.FindOneById(payload.ID)
	if err != nil {
		jsonErr := http.BadRequestError(err)
		return &http.APIResponse{APIError: jsonErr}
	}

	userDTO := a.UserReqDataMapper.ToDalEntity(user)
	jsonOk := http.SuccessResponse(userDTO, "User Found Successfully")
	return &http.APIResponse{APISuccess: jsonOk}
}

func (a *UserApplicationService) CreateUser(payload *dto.CreateUserDTO) *http.APIResponse {
	userDTO := &dto.UserDTO{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Address:   payload.Address,
	}
	user := a.UserReqDataMapper.ToDomainEntity(userDTO)
	if err := a.userService.CreateUser(user); err != nil {
		jsonErr := http.BadRequestError(err)
		return &http.APIResponse{APIError: jsonErr}
	}

	jsonOk := http.SuccessResponse(&dto.AcknowledgeDTO{
		Action:    "create_user",
		IsSuccess: true,
	}, "User Created Successfully")
	return &http.APIResponse{APISuccess: jsonOk}
}
