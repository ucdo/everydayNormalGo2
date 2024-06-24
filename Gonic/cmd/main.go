package main

import (
	"Gonic/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Todo List API built with Gin",
		})
	})

	router.GET("/ping", handler.CreateTodoTable)
	router.GET("list", handler.TodoList)

	router.Run()
}
