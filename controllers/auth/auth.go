package auth

import (
	"encoding/json"
	"net/http"
	"quik/logger"
	"quik/types"
	"quik/utils"

	"github.com/gin-gonic/gin"
)

// This is just added to help get a sign in token
// nothing serious happening here
// it can be improved on :)
func Login(c *gin.Context) {
	var data types.LoginRequest
	input := json.NewDecoder(c.Request.Body)

	if err := input.Decode(&data); err != nil {
		logger.Error(
			"error while unmarshaling credit data",
			err, map[string]interface{}{
				"method": c.Request.Method,
				"url":    c.Request.URL.String(),
				"data":   c.Request.Body,
			})

		utils.InvalidRequestBodyAPIError(c)
		return
	}

	if ok, errValidate := utils.StructValidateHelper(data); ok {
		utils.RequestValidationAPIError(c, errValidate[0])
		return
	}

	token, err := utils.CreateToken(data.User)
	if err != nil {
		utils.InternalServerErrorAPIError(c, "something went wrong")
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
