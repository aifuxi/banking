package errs

import "net/http"

type AppErr struct {
	Code    int
	Message string
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
