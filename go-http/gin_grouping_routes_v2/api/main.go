package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	//Authen Group
	authenAPI := router.Group("/authen")
	{
		// http://localhost:85/authen/login
		authenAPI.GET("/login", Login)
		authenAPI.GET("/register", Register)
		authenAPI.GET("/profile", func(c *gin.Context) {
			c.String(http.StatusOK, "profile")
		})
	}

	stockAPI := router.Group("/stock")
	{
		// http://localhost:85/stock/list
		stockAPI.GET("/list", ListProduct)
		stockAPI.GET("/create", CreateProduct)
	}
}
