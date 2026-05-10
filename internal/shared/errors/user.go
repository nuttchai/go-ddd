package errors

import (
	"errors"

	constant "github.com/nuttchai/go-ddd/internal/shared/constants"
)

var (
	ErrEmailAlreadyExisted = errors.New(constant.EmailAlreadyExisted)
	ErrInvalidCreatedUser  = errors.New(constant.InvalidCreatedUser)
	ErrInvalidUpdatedUser  = errors.New(constant.InvalidUpdatedUser)
)
