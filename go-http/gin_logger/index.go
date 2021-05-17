package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	count := 0

	runningDir, _ := os.Getwd()

	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/logs/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/logs/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile

	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("root"))
	})

	r.GET("/error", func(c *gin.Context) {
		count = count + 1
		errlogfile.WriteString(fmt.Sprintf("Error Count : %d \n", count))
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("error"))
	})

	r.GET("/profile", func(c *gin.Context) {
		count = count + 1
		accesslogfile.WriteString(fmt.Sprintf("Count : %d \n", count))
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("profile"))
	})

	r.Run(":85")

}
