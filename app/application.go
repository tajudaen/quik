package app

import (
	"fmt"
	"log"

	"quik/config"
	"quik/middlewares"
	"quik/providers/logger"
	"quik/utils/errors"

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

	router.NoRoute(func(c *gin.Context) {
		errors.NotFoundAPIError(c)
	})

	if err := router.Run(fmt.Sprintf(":%s", config.C.Port)); err != nil {
		log.Fatalf("Can't run the server, err: %v\n", err)
	}
}
