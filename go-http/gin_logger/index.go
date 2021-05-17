package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Use for disable  Log output coloring
	//gin.DisableConsoleColor()
	count := 0

	runningDir, _ := os.Getwd()

	errlogfile, _ := os.OpenFile(fmt.Sprintf("%s/logs/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/logs/gin_access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	gin.DefaultErrorWriter = errlogfile
	gin.DefaultWriter = accesslogfile

	// Basic Logger
	//r.Use(gin.Logger())

	// Logger with format
	/*
		r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			return fmt.Sprintf("%s [%s]  \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		}))
	*/

	// Disable logging for "/profile"
	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/profile"))

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
		accesslogfile.WriteString(fmt.Sprintf("Count : %d ", count))
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("profile"))
	})

	r.Run(":85")

}
