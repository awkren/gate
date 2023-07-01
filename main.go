package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"gate/handlers"
	"gate/internal"
	"gate/middlewares"
	"gate/models"
	"gate/utils"
)

func main() {
	// Load environment variables from .env
	err := utils.LoadEnvVariables()
	if err != nil {
		fmt.Println("Error loading .env file", err)
		return
	}

	// Initialize the permission map
	handlers.PermissionMap = make(map[models.Request]bool)

	// Retrieve the external API address and endpoint from .env vars
	handlers.ExternalAPI = os.Getenv("EXTERNAL_API_ADDRESS")
	handlers.Endpoint = os.Getenv("EXTERNAL_API_ENDPOINT")
	handlers.UserEndpoint = os.Getenv("USER_ENDPOINT")

	// Create a new Gin router
	router := gin.Default()

	// apply the ip filtering to all routes
	router.Use(middlewares.IPFilterMiddleware)

	// Define your API routes
	router.Use(middlewares.CheckPermissionMiddleware)

	// create a group for the users endpoint
	usersGroup := router.Group(handlers.UserEndpoint)
	{
		usersGroup.GET("", handlers.HandleUserRequest)
	}

	// handle requests to other endpoints
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Endpoint not found :p",
		})
	})

	// Start the server
	router.Run(":8080")
}

func init() {
	handlers.Requests = make(chan models.Request)

	go internal.ReadApprovalInput()
}
