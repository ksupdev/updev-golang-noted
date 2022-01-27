package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"labs.test/opl-jager02/micro_serv01/handler"
	"labs.test/opl-jager02/share/adc_trace"
	"labs.test/opl-jager02/share/middleware"
)

func main() {
	servName := "micro-service01"
	log.Printf("%v", servName)

	jaegerURL := "http://localhost:14268/api/traces"
	if env, ok := os.LookupEnv("JAEGER_URL"); ok {
		jaegerURL = env
	}

	adc_trace.SetupOtel(jaegerURL, servName)
	hdl := handler.NewMainHandler(otel.Tracer(servName))

	r := gin.Default()
	r.Use(middleware.Middleware(servName))

	r.POST("/serv-01", hdl.Handler)

	// http://localhost:8021/serv-01
	r.Run(":8021")
}
