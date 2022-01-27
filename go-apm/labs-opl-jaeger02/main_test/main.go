package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// servName := "main_service"
	// hdl := handler.NewMainHandler(service, adc_tracer)

	r := gin.Default()

	r.POST("/serv-main", Handler)
	// r.ServeHTTP()
	// otelhttp.NewHandler(r, servName)
	// r.GET("/serv-main", otelhttp.NewHandler(hdl.Handler, servName))

	// http://localhost:8010/serv-main

	r.Run(":8099")
}

func Handler(c *gin.Context) {
	type Resp struct {
		ServName string `json:"appname"`
		Message  string `json:"version"`
	}

	// Implement logic

	reqBody := ServRequest{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		message := fmt.Sprintf("request body incorrect format ,%v", err)
		c.JSON(http.StatusUnprocessableEntity, Resp{ServName: "microservice 01", Message: message})
		return
	}

	c.JSON(http.StatusOK, Resp{ServName: "main-service", Message: "Say hi " + reqBody.WhoRequest})
}

type ServRequest struct {
	WhoRequest string `json:"who_request"`
	Data       string `json:"data"`
}
