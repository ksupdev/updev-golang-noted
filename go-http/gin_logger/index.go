package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	runningDir, _ := os.Getwd()

	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/logs/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/logs/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile

	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("root"))
	})

	r.GET("/profile", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("profile"))
	})

	r.Run(":85")

}
