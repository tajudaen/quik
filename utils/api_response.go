package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIError(c *gin.Context, err *RestErr) {
	c.JSON(err.Status, err)
}

func APISuccess(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func NotFoundAPIError(c *gin.Context) {
	err := NewErrorResponse("resource not available", http.StatusNotFound, "not_found")

	APIError(c, err)
}

func InternalServerErrorAPIError(c *gin.Context, e string) {
	err := NewErrorResponse("something went wrong", http.StatusInternalServerError, "internal_server_error")

	APIError(c, err)
}

func InvalidRequestBodyAPIError(c *gin.Context) {
	err := NewErrorResponse("invalid request body provided", http.StatusBadRequest, "bad_request")

	APIError(c, err)
}

func RequestValidationAPIError(c *gin.Context, e string) {
	err := NewErrorResponse(fmt.Sprintf("invalid request body provided. Error: %v", e), http.StatusBadRequest, "bad_request")

	APIError(c, err)
}
