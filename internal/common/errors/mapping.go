package errors

import (
	"net/http"

	chttp "github.com/nuttchai/go-ddd/common/http"
)

const (
	CodeUserInvalidCreated chttp.Code = "2001"
	CodeUserInvalidUpdated chttp.Code = "2002"
	CodeUserEmailExists    chttp.Code = "2003"
)

var ErrorMappings = []chttp.ErrorMapping{
	{Sentinel: ErrEmailAlreadyExisted, Status: http.StatusBadRequest, Code: CodeUserEmailExists},
	{Sentinel: ErrInvalidCreatedUser, Status: http.StatusBadRequest, Code: CodeUserInvalidCreated},
	{Sentinel: ErrInvalidUpdatedUser, Status: http.StatusBadRequest, Code: CodeUserInvalidUpdated},
}
