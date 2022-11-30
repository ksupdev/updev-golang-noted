package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	hdl "puza.dev/go-labs-swagger/handler"
)

func main() {
	fmt.Println("test")

	barHdl := hdl.NewBar()

	router := initWebServer()
	contextRootV1 := router.Group("api/v1")
	{
		contextRootV1.GET("/bar/:id", barHdl.GetById)
	}

	err := router.Run(":8081")
	if err != nil {
		panic(err)
	}

}

func initWebServer() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))
	// logf.Info("[Startup] Webserver success ")
	return router
}
