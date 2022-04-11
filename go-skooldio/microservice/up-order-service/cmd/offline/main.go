package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"updev.labs/up-order-service/order"
	"updev.labs/up-order-service/router"
	"updev.labs/up-order-service/store"
)

func init() {
	err := godotenv.Load("offline.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

func main() {

	fmt.Println(os.Getenv("FILTER_CHANNEL"))
	r := router.NewRouter()
	handler := order.NewHandler(os.Getenv("FILTER_CHANNEL"), store.NewMariaDBStoreMock(os.Getenv("DSN")))
	r.POST("/api/v1/orders", handler.Order)

	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	stop()

	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}
