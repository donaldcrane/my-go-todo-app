package main

import (
	"github.com/gin-gonic/gin"

	"github.com/donaldcrane/my-go-todo-api/config"
	"github.com/donaldcrane/my-go-todo-api/models"
)

func main() {

    config.ConnectDatabase()

    config.DB.AutoMigrate(&models.Todo{})
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
			"language": "Go!",
			"framework": "Gin",
		})
	})

	router.Run(":3000")
}