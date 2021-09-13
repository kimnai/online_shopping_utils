package error

import (
	"fmt"
	"net/http"
)

type ErrorInterface interface {
	Message() string
	Status() int
}

type error struct {
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
}

func (e error) Message() string {
	return fmt.Sprintf("Error Status: %d,  Error Message: %s", e.ErrorStatus, e.ErrorMessage)
}

func (e error) Status() int {
	return e.ErrorStatus
}

func NewBadRequestError(message string) ErrorInterface {
	return error{
		ErrorMessage: message,
		ErrorStatus:  http.StatusBadRequest,
	}
}

func NewInternalServerError(message string) ErrorInterface {
	return error{
		ErrorMessage: message,
		ErrorStatus:  http.StatusInternalServerError,
	}
}

func NewUnauthorizedError(message string) ErrorInterface {
	return error{
		ErrorMessage: message,
		ErrorStatus:  http.StatusUnauthorized,
	}
}
