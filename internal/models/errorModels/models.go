package errorModels

import (
	"time"
)

type ServiceError struct {
	Time        string `json:"time"`
	Description string `json:"description"`
	Error       any    `json:"error"`
}

func NewCustomServiceError(description string, err error) *ServiceError {
	return &ServiceError{Time: time.Now().Format(time.RFC3339), Description: description, Error: err.Error()}
}

func NewInternalServiceError(err error) *ServiceError {
	return &ServiceError{Time: time.Now().Format(time.RFC3339), Description: "Internal Service Error", Error: err.Error()}
}

func NewValidatingError(err error) *ServiceError {
	return &ServiceError{Time: time.Now().Format(time.RFC3339), Description: "Bad Input Request", Error: err.Error()}
}
