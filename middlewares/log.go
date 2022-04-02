package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogsMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("router middleware",map[string]interface{}{
			"method": c.Request.Method,
			"url": c.Request.URL.String(),
			"data": c.Request.Body,
		},
		)
		c.Next()
	}
}
