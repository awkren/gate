package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Method string
	Path   string
}

var (
	permissionMap  map[Request]bool
	permissionLock sync.RWMutex
	requests       chan Request
)

func main() {
	// Initialize the permission map
	permissionMap = make(map[Request]bool)

	// Create a new Gin router
	router := gin.Default()

	// Define your API routes
	router.Use(checkPermissionMiddleware)
	router.GET("/users", handleUserRequest)

	// Start the server
	router.Run(":8080")
}

func checkPermissionMiddleware(c *gin.Context) {
	req := Request{
		Method: c.Request.Method,
		Path:   c.Request.URL.Path,
	}

	// Send the request to the permission check channel
	requests <- req

	permissionLock.RLock()
	permission, exists := permissionMap[req]
	permissionLock.RUnlock()

	if exists && permission {
		// Continue processing the request
		c.Next()
	} else {
		// Respond with an error message or denial
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Permission denied",
		})
		c.Abort()
	}
}

func handleUserRequest(c *gin.Context) {
	// Handle the user request
	c.JSON(http.StatusOK, gin.H{
		"message": "User request handled successfully",
	})
}

func readApprovalInput() {
	reader := bufio.NewReader(os.Stdin)

	for req := range requests {
		fmt.Printf("Received request: %s %s\n", req.Method, req.Path)
		fmt.Print("Do you want to grant permission for this request? (Y/N): ")

		decision, _ := reader.ReadString('\n')
		decision = strings.TrimSpace(strings.ToLower(decision))

		permission := decision == "y" || decision == "yes"

		permissionLock.Lock()
		permissionMap[req] = permission
		permissionLock.Unlock()

		fmt.Printf("Permission for request %s %s set to %v\n", req.Method, req.Path, permission)
	}
}

func init() {
	requests = make(chan Request)

	go readApprovalInput()
}
