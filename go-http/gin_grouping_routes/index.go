package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.String(http.StatusOK, "Login")
}

func register(c *gin.Context) {
	c.String(http.StatusOK, "Register")
}

func listProduct(c *gin.Context) {
	c.String(http.StatusOK, "List Product")
}

func createProduct(c *gin.Context) {
	c.String(http.StatusOK, "Create Product")
}

func main() {
	router := gin.Default()

	//Authen Group
	authenAPI := router.Group("/authen")
	{
		// http://localhost:85/authen/login
		authenAPI.GET("/login", login)
		authenAPI.GET("/register", register)
		authenAPI.GET("/profile", func(c *gin.Context) {
			c.String(http.StatusOK, "profile")
		})
	}

	stockAPI := router.Group("/stock")
	{
		// http://localhost:85/stock/list
		stockAPI.GET("/list", listProduct)
		stockAPI.GET("/create", createProduct)
	}

	router.Run(":85")

}
