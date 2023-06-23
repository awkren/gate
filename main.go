package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	permissionGranted bool
)

func main() {
	// Set initial permission status to false
	permissionGranted = true

	// Create a new Gin router
	router := gin.Default()

	// Define your API routes
	router.GET("/users", handleUserRequest)
	router.GET("/admin/approve", handleAdminApproval)

	// Start the server
	router.Run(":8080")
}

func handleUserRequest(c *gin.Context) {
	if permissionGranted {
		// Forward the request to the Node.js API
		http.Redirect(
			c.Writer,
			c.Request,
			"http://localhost:3000/users",
			http.StatusTemporaryRedirect,
		)
	} else {
		// Respond with an error message or denial
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Permission denied",
		})
	}
}

func handleAdminApproval(c *gin.Context) {
	// Set the permission status to true upon manual approval
	permissionGranted = true

	fmt.Println("Permission granted manually by the administrator")

	c.JSON(http.StatusOK, gin.H{
		"message": "Permission granted",
	})
}
