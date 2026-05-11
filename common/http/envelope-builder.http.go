package http

import (
	"errors"
	nethttp "net/http"
)

type ErrorMapping struct {
	Sentinel error
	Status   int
	Code     Code
}

type EnvelopeBuilder struct {
	mappings []ErrorMapping
}

func NewEnvelopeBuilder(mappings []ErrorMapping) *EnvelopeBuilder {
	return &EnvelopeBuilder{mappings: mappings}
}

func (b *EnvelopeBuilder) Build(data any, err error) *EnvelopeResult {
	if err != nil {
		status, code := b.mapError(err)
		return Err(err.Error(), status, code)
	}
	return OK(data, nethttp.StatusOK)
}

func (b *EnvelopeBuilder) mapError(err error) (int, Code) {
	for _, m := range b.mappings {
		if errors.Is(err, m.Sentinel) {
			return m.Status, m.Code
		}
	}
	return nethttp.StatusInternalServerError, CodeExceptionError
}
