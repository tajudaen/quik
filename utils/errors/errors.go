package errors

import (
	"net/http"
	"quik/utils"

	"github.com/gin-gonic/gin"
)

func APIError(c *gin.Context, message string, status int, eMessage string) {
	err := utils.NewErrorResponse(message, status, eMessage)
	c.JSON(status, err)
}

func NotFoundAPIError(c *gin.Context) {
	APIError(c, "resource not available", http.StatusNotFound, "not_found")
}

func InternalServerErrorAPIError(c *gin.Context, e string) {
	message := "something went wrong"
	status := http.StatusInternalServerError
	eMessage := "internal_server_error"

	APIError(c, message, status, eMessage)
}
