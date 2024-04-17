package errs

import "net/http"

type AppErr struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e AppErr) AsMessage() *AppErr {
	return &AppErr{Message: e.Message}
}

func NewNotFoundErr(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectErr(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewValidationErr(message string) *AppErr {
	return &AppErr{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
