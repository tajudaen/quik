package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"quik/config"
	"quik/middlewares"
	"quik/providers/logger"
)

var (
	router = gin.New()
)

func StartApplication() {

	router.Use(middlewares.LogsMiddleware(logger.Log))
	router.Run(fmt.Sprintf(":%s", config.C.Port))
}
