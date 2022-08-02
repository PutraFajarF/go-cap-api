package errs

import (
	"net/http"
)

type AppErr struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e AppErr) AsMessage() *AppErr {
	return &AppErr{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewBadRequestError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewValidationError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}

func NewAuthenticationError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}
<<<<<<< HEAD
=======

func NewForbiddenError(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusForbidden,
		Message: message,
	}
}
>>>>>>> c10f03cd41226d7095285e2635a29eef01f73db9
