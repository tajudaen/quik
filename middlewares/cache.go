package middlewares

import (
	"net/http"
	"quik/providers/redis"

	"github.com/gin-gonic/gin"
)

func CacheBalance() gin.HandlerFunc {
	return func(c *gin.Context) {
		cache := &redis.Redis{}
		balance, err := cache.Get(c.Param("wallet_id"))
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"balance": balance})
			c.AbortWithStatus(http.StatusOK)
		}

		c.Next()
	}
}
