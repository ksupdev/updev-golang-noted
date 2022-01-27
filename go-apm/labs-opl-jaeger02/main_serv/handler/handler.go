package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"labs.test/opl-jager02/share/util"
)

// Handler section

func NewMainHandler(service *Service, tracer trace.Tracer) *MainHandler {
	return &MainHandler{service: service, Tracer: tracer}
}

type MainHandler struct {
	service *Service
	Tracer  trace.Tracer
}

func (mhdl *MainHandler) Handler(c *gin.Context) {
	type Resp struct {
		ServName string `json:"appname"`
		Message  string `json:"version"`
	}
	// ctx := c.Request.Context()
	// span := trace.SpanFromContext(ctx)
	// ctx := c.Request.Context()
	ctx, span := mhdl.Tracer.Start(c.Request.Context(), "handle main_service")
	defer span.End()

	util.DoSomething()

	servInput := ServRequest{Data: "first data"}

	servOut, err := mhdl.service.RequestToOtherServiceNative(ctx, servInput)
	if err != nil {
		message := fmt.Sprintf("error request to  ,%v", err)
		c.JSON(http.StatusUnprocessableEntity, Resp{ServName: "main-service", Message: message})
		return
	}

	c.JSON(http.StatusOK, Resp{ServName: "main-service", Message: servOut.Message})
}
