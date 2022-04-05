package middlewares

import (
	"fmt"
	"net/http"
	"quik/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if bearerToken := c.Request.Header.Get("Authorization"); bearerToken != "" {
			_, err := utils.ExtractTokenMetadata(bearerToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"err": fmt.Sprintf("%s", err)})
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"err": "no token provided"})
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Next()
	}
}
