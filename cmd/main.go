package main

import (
	"github.com/gin-gonic/gin"

	"github.com/donaldcrane/my-go-todo-api/config"
	"github.com/donaldcrane/my-go-todo-api/handlers"
	"github.com/donaldcrane/my-go-todo-api/models"
	"github.com/donaldcrane/my-go-todo-api/repositories"
	"github.com/donaldcrane/my-go-todo-api/services"
)

func main() {

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Todo{})

	// Repository

	todoRepository := repositories.NewTodoRepository(
		config.DB,
	)

	// Service

	todoService := services.NewTodoService(
		todoRepository,
	)

	// Handler

	todoHandler := handlers.NewTodoHandler(
		todoService,
	)

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Donald is learning Go!",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"language":  "Go!",
			"framework": "Gin",
		})
	})

	router.POST("/todos", todoHandler.CreateTodo)
	router.GET("/todos", todoHandler.GetTodos)
	router.GET("/todos/:id", todoHandler.GetTodo)
	router.GET("/todos/completed", todoHandler.GetCompleted)
	router.PUT("/todos/:id", todoHandler.UpdateTodo)
	router.PUT("/todos/:id/completed", todoHandler.CompleteTodo)
	router.DELETE("/todos/:id", todoHandler.DeleteTodo)
	router.Run(":3000")
}
