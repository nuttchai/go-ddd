package errors

import (
	"net/http"

	code "github.com/nuttchai/go-ddd/common/http"
)

const (
	CodeUserInvalidCreated code.Code = "2001"
	CodeUserInvalidUpdated code.Code = "2002"
	CodeUserEmailExists    code.Code = "2003"
)

type errorMapping struct {
	Sentinel error
	Status   int
	Code     code.Code
}

var ErrorMappings = []errorMapping{
	{ErrEmailAlreadyExisted, http.StatusBadRequest, CodeUserEmailExists},
	{ErrInvalidCreatedUser, http.StatusBadRequest, CodeUserInvalidCreated},
	{ErrInvalidUpdatedUser, http.StatusBadRequest, CodeUserInvalidUpdated},
}
