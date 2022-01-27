package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func handleBookRequest(c *gin.Context) {
	from, to := c.Param("from"), c.Param("to")
	vichicle := c.Param("vichicle")
	c.JSON(http.StatusOK, gin.H{"result": "ok", "from": from, "to": to, "vichicle": vichicle})

}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		// Response send plain text
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Hi this root service"))
	})

	r.GET("/profile", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("Hi this profile page"))
	})

	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "username": username, "password": password})
		//gin.H => HashMap with gin
	})

	r.GET("/book/:from/:to/:vichicle", handleBookRequest)

	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		/*
			c.ShouldBind use for mapping post form data to Struct
			if c.ShouldBind(&form) != nill when you pass value miss match with struct or condition
		*/
		if c.ShouldBind(&form) == nil {

			if form.Username == "admin" && form.Password == "1234" {
				msg := fmt.Sprintf("You are logged with: %s %s", form.Username, form.Password)
				c.JSON(http.StatusOK, gin.H{"status": msg})
				//c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				fmt.Printf(" user : %s , pass : %s", form.Username, form.Password)
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}

		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unable to bind"})
		}

	})

	// config port and start server
	r.Run(":9999") //http://localhost:9999/

}
