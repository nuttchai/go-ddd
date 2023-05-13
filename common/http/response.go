package api

type APIResponse struct {
	APISuccess *APISuccess
	APIError   *APIError
}

func (r *APIResponse) Value() any {
	if r.APIError != nil {
		return *r.APIError
	}
	return *r.APISuccess
}

func (r *APIResponse) Status() int {
	if r.APIError != nil {
		return r.APIError.Status
	}
	return r.APISuccess.Status
}
