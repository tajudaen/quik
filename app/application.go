package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"quik/config"
)

var (
	router = gin.Default()
)

func StartApplication() {
	router.Run(fmt.Sprintf(":%s", config.C.Port))
}
