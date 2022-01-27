package main

import (
	"updev/gin-grp-route-v2/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api.Setup(router)
	router.Run(":85")

}
