package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/users", handleUserRequest)

	router.Run(":8080")
}

func handleUserRequest(c *gin.Context) {
	permissionGranted := checkPermission()

	if permissionGranted {
		http.Redirect(
			c.Writer,
			c.Request,
			"http://localhost:3000/users",
			http.StatusTemporaryRedirect,
		)
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"Error": "Permission denied",
		})
	}
}

func checkPermission() bool {
	return false
}
