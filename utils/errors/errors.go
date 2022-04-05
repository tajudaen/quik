package errors

import (
	"net/http"
	"quik/utils"
)

func NewInternalServerError(message string) *utils.RestErr {
	return &utils.RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Err:     "internal_server_error",
	}
}

func NewNotFoundError(message string) *utils.RestErr {
	return &utils.RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Err:     "not_found",
	}
}

func NewInsufficientError(message string) *utils.RestErr {
	return &utils.RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Err:     "bad_request",
	}
}
