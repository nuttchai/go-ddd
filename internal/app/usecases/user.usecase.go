package usecase

import (
	"errors"

	http "github.com/nuttchai/go-ddd/common/http"
	irepository "github.com/nuttchai/go-ddd/internal/app/repositories"
	httpdtos "github.com/nuttchai/go-ddd/internal/http/dtos"
	constant "github.com/nuttchai/go-ddd/internal/shared/constants"
)

type UserUsecase struct {
	userRepo irepository.IUserRepository
}

func NewUserUsecase(userRepo irepository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (a *UserUsecase) FindUserById(payload *httpdtos.FindUserByIdDTO) *http.APIResponse {
	user, err := a.userRepo.FindOneById(payload.ID)
	if err != nil {
		jsonErr := http.BadRequestError(err)
		return &http.APIResponse{APIError: jsonErr}
	}

	jsonOk := http.SuccessResponse(userToResponseDTO(user), "User Found Successfully")
	return &http.APIResponse{APISuccess: jsonOk}
}

func (a *UserUsecase) CreateUser(payload *httpdtos.CreateUserDTO) *http.APIResponse {
	user := userFromCreateDTO(payload)
	if isUserValid := user.IsUserValid(); !isUserValid {
		jsonErr := http.BadRequestError(errors.New(constant.InvalidCreatedUser))
		return &http.APIResponse{APIError: jsonErr}
	}
	if userDb, err := a.userRepo.FindOneByEmail(user.Email); userDb != nil || err != nil {
		if err != nil {
			jsonErr := http.BadRequestError(err)
			return &http.APIResponse{APIError: jsonErr}
		}
		jsonErr := http.BadRequestError(errors.New(constant.EmailAlreadyExisted))
		return &http.APIResponse{APIError: jsonErr}
	}
	if err := a.userRepo.Save(user); err != nil {
		jsonErr := http.BadRequestError(err)
		return &http.APIResponse{APIError: jsonErr}
	}

	jsonOk := http.SuccessResponse(&httpdtos.AcknowledgeDTO{
		Action:    "create_user",
		IsSuccess: true,
	}, "User Created Successfully")

	return &http.APIResponse{APISuccess: jsonOk}
}

func (a *UserUsecase) UpdateUser(payload *httpdtos.UpdateUserDTO) *http.APIResponse {
	user := userFromUpdateDTO(payload)
	if isUserValid := user.IsUserValid(); !isUserValid {
		jsonErr := http.BadRequestError(errors.New(constant.InvalidUpdatedUser))
		return &http.APIResponse{APIError: jsonErr}
	}
	if err := a.userRepo.UpdateUser(user); err != nil {
		jsonErr := http.BadRequestError(err)
		return &http.APIResponse{APIError: jsonErr}
	}

	jsonOk := http.SuccessResponse(&httpdtos.AcknowledgeDTO{
		Action:    "update_user",
		IsSuccess: true,
	}, "User Updated Successfully")

	return &http.APIResponse{APISuccess: jsonOk}
}
