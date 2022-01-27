package main

import (
	"log"
	"os"

	// gintrace "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"labs.test/opl-jager02/main_serv/handler"
	"labs.test/opl-jager02/share/adc_trace"
	"labs.test/opl-jager02/share/middleware"
)

func main() {
	servName := "main_service"
	log.Printf("%v", servName)

	jaegerURL := "http://localhost:14268/api/traces"
	if env, ok := os.LookupEnv("JAEGER_URL"); ok {
		jaegerURL = env
	}

	adc_trace.SetupOtel(jaegerURL, servName)

	adc_tracer := otel.Tracer(servName)

	service := handler.NewService(servName, adc_tracer)
	hdl := handler.NewMainHandler(service, adc_tracer)

	r := gin.Default()
	r.Use(middleware.Middleware(servName))

	r.GET("/serv-main", hdl.Handler)
	// r.ServeHTTP()
	// otelhttp.NewHandler(r, servName)
	// r.GET("/serv-main", otelhttp.NewHandler(hdl.Handler, servName))

	// http://localhost:8010/serv-main

	r.Run(":8010")
}
