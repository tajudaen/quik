package utils

import (
	"fmt"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Err     string `json:"error"`
}

type RestSuccess struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func (r *RestErr) Error() string {
	return fmt.Sprintf("message:%s status %d: err %v", r.Message, r.Status, r.Err)
}

func NewErrorResponse(message string, status int, err string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Err:     err,
	}
}

func NewSuccessResponse(message string, status int, data interface{}) *RestSuccess {
	return &RestSuccess{
		Message: message,
		Status:  status,
		Data:    data,
	}
}
