package models

import "net/http"

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}

func NewAPIError(code int, message string, details ...string) *APIError {
	err := &APIError{
		Code:    code,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

var (
	ErrUserNotFound     = NewAPIError(http.StatusNotFound, "User not found")
	ErrInvalidInput     = NewAPIError(http.StatusBadRequest, "Invalid input")
	ErrInternalServer   = NewAPIError(http.StatusInternalServerError, "Internal server error")
	ErrDatabaseError    = NewAPIError(http.StatusInternalServerError, "Database error")
	ErrValidationFailed = NewAPIError(http.StatusBadRequest, "Validation failed")
)
