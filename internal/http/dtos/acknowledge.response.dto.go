package dtos

type AcknowledgeDTO struct {
	Action    string `json:"action"`
	IsSuccess bool   `json:"is_success"`
}
