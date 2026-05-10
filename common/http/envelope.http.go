package http

import nethttp "net/http"

type Code string

const (
	CodeOK              Code = "0000"
	CodeExceptionError  Code = "1000"
	CodeBadRequestError Code = "2000"
)

type Envelope struct {
	Code  Code   `json:"code"`
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

type EnvelopeResult struct {
	Status   int
	Envelope Envelope
}

func OK(data any, status int) *EnvelopeResult {
	env := Envelope{Code: CodeOK}
	if data != nil {
		env.Data = data
	}
	return &EnvelopeResult{Status: status, Envelope: env}
}

func Err(message string, status int, code Code) *EnvelopeResult {
	return &EnvelopeResult{
		Status: status,
		Envelope: Envelope{
			Code:  code,
			Error: message,
		},
	}
}

func ErrWithData(message string, status int, code Code, data any) *EnvelopeResult {
	return &EnvelopeResult{
		Status: status,
		Envelope: Envelope{
			Code:  code,
			Error: message,
			Data:  data,
		},
	}
}

func MapErr(err error) *EnvelopeResult {
	if err == nil {
		return nil
	}
	return Err(err.Error(), nethttp.StatusInternalServerError, CodeExceptionError)
}
