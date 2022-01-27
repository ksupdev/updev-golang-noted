package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}

		msg.Name = "Karoon"
		msg.Message = "Hey"
		msg.Number = 123

		// Note that msg.Name becomes "user" in the JSON
		// Will output : {"user":"Karoon","Message":"Hey","Number":123}

		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someString", func(c *gin.Context) {
		c.String(http.StatusOK, "gin.H{message: hey, status: http.StatusOK}")
	})

	r.Run(":85")

}
