package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"labs.test/opl-jager02/share/util"
)

// Handler section

func NewMainHandler(tracer trace.Tracer) *MainHandler {
	return &MainHandler{Tracer: tracer}
}

type MainHandler struct {
	Tracer trace.Tracer
}

func (mhdl *MainHandler) Handler(c *gin.Context) {

	_, span := mhdl.Tracer.Start(c.Request.Context(), "handle micro service 01")
	defer span.End()

	if c.GetHeader("content-type") != "application/json" {
		message := fmt.Sprintf("error contect-type request ,%v", c.GetHeader("content-type"))
		span.RecordError(fmt.Errorf(message))
		span.SetStatus(codes.Error, message)

		c.JSON(http.StatusUnprocessableEntity, Resp{ServName: "microservice 01", Message: message, Status: "fail"})
		return
	}

	reqBody := ServRequest{}
	err := c.ShouldBind(&reqBody)
	if err != nil {
		message := fmt.Sprintf("request body incorrect format ,%v", err)
		c.JSON(http.StatusUnprocessableEntity, Resp{ServName: "microservice 01", Message: message, Status: "fail"})
		return
	}

	util.DoSomething()
	// modify for response
	outbout := fmt.Sprintf("Microservice 01 recieve %v data form ,%v", reqBody.Data, reqBody.WhoRequest)

	c.JSON(http.StatusOK, Resp{ServName: "microservice 01", Message: outbout, Status: "success"})
}

type Resp struct {
	ServName string `json:"service_name"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

type RespErr struct {
	ServName string `json:"service_name"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}
