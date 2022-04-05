package app

import (
	"fmt"
	"log"

	"quik/config"
	"quik/logger"
	"quik/middlewares"
	"quik/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.New()
)

func StartApplication() {
	router.Use(gin.Recovery())

	router.Use(middlewares.LogsMiddleware(logger.Log))
	router.Use(cors.Default())

	registerRoutes()

	router.NoRoute(func(c *gin.Context) {
		utils.NotFoundAPIError(c)
	})

	if err := router.Run(fmt.Sprintf(":%s", config.C.Port)); err != nil {
		log.Fatalf("Can't run the server, err: %v\n", err)
	}
}
