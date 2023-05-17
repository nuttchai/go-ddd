package http

import (
	"net/http"
)

type APISuccess struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

func SuccessResponse(result any, message ...string) *APISuccess {
	msg := defaultAPISuccessMsg
	if len(message) > 0 {
		msg = message[0]
	}

	return &APISuccess{
		Status:  http.StatusOK,
		Message: msg,
		Result:  result,
	}
}
