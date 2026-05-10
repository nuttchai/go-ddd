package helpers

import (
	"errors"
	nethttp "net/http"

	chttp "github.com/nuttchai/go-ddd/common/http"
	shared "github.com/nuttchai/go-ddd/internal/shared/errors"
)

func mapError(err error) (status int, code chttp.Code) {
	for _, m := range shared.ErrorMappings {
		if errors.Is(err, m.Sentinel) {
			return m.Status, m.Code
		}
	}
	return nethttp.StatusInternalServerError, chttp.CodeExceptionError
}

func BuildEnvelope(data any, err error) *chttp.EnvelopeResult {
	if err != nil {
		status, code := mapError(err)
		return chttp.Err(err.Error(), status, code)
	}
	return chttp.OK(data, nethttp.StatusOK)
}
