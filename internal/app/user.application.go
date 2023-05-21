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
	userReqDataMapper cmapper.IDataMapper[entity.User, dto.UserDTO]
}

func NewUserApplicationService(userService service.IUserService, userReqDataMapper cmapper.IDataMapper[entity.User, dto.UserDTO]) IUserApplicationService {
	return &UserApplicationService{
		userService:       userService,
		userReqDataMapper: userReqDataMapper,
	}
}

func (a *UserApplicationService) FindUserById(payload *dto.FindUserByIdDTO) *http.APIResponse {
	user, err := a.userService.FindOneById(payload.ID)
	if err != nil {
		jsonErr := http.BadRequestError(err)
		return &http.APIResponse{APIError: jsonErr}
	}

	userDTO := a.userReqDataMapper.ToDalEntity(user)
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
	user := a.userReqDataMapper.ToDomainEntity(userDTO)
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

func (a *UserApplicationService) UpdateUser(payload *dto.UpdateUserDTO) *http.APIResponse {
	userDTO := &dto.UserDTO{
		ID:        payload.ID,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Address:   payload.Address,
	}
	user := a.userReqDataMapper.ToDomainEntity(userDTO)
	if err := a.userService.UpdateUser(user); err != nil {
		jsonErr := http.BadRequestError(err)
		return &http.APIResponse{APIError: jsonErr}
	}

	jsonOk := http.SuccessResponse(&dto.AcknowledgeDTO{
		Action:    "update_user",
		IsSuccess: true,
	}, "User Updated Successfully")

	return &http.APIResponse{APISuccess: jsonOk}
}
