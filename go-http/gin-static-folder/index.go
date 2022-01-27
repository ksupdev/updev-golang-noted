package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//http://localhost:85/assets/ => Config directory allow access from external system
	r.Static("/assets", "./public")
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hi")
	})

	r.Run(":85")
}
