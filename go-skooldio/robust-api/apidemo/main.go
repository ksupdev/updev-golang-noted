package main

import (
	"github.com/gin-gonic/gin"
	"updev.co.th/apidemo/todo"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handler := todo.NewTodoHandler()
	r.POST("/todos", handler.NewTask)

	r.Run()
}

// https://github.com/go-playground/validator
