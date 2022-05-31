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

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/updev/todoapi/todo"
)

var (
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

func main() {
	// err = godotenv.Load("local.env")
	// if err != nil {
	// 	log.Printf("please consider environment variables: %s\n", err)
	// }

	// db, err := gorm.Open(sqlite.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// if err := db.AutoMigrate(&todo.Todo{}); err != nil {
	// 	log.Println("auto migrate db", err)
	// }

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
		"TransactionID",
	}
	r.Use(cors.New(config))

	r.GET("/healthz", func(c *gin.Context) {
		c.Status(200)
	})
	r.GET("/limitz", limitedHandler)
	r.GET("/x", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"buildcommit": buildcommit,
			"buildtime":   buildtime,
		})
	})

	gormStore := todo.NewGormStoreFake()

	handler := todo.NewTodoHandler(gormStore)
	r.POST("/todos", todo.NewGinHandler(handler.NewTask))
	r.GET("/todos", todo.NewGinHandler(handler.List))
	r.DELETE("/todos/:id", todo.NewGinHandler(handler.Remove))
	// r.DELETE("/todos/:id", handler.Remove)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}

var limiter = rate.NewLimiter(5, 5)

func limitedHandler(c *gin.Context) {
	if !limiter.Allow() {
		c.AbortWithStatus(http.StatusTooManyRequests)
		return
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
