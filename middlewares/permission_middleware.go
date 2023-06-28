package middlewares

import (
	"fmt"
	"gate/handlers"
	"gate/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CheckPermissionMiddleware(c *gin.Context) {
	req := models.Request{
		Method: c.Request.Method,
		Path:   c.Request.URL.Path,
	}

	// log the request into a file
	err := logRequest(req)
	if err != nil {
		fmt.Println("Error man kkkkkk")
	}

	// Send the request to the permission check channel
	handlers.Requests <- req

	handlers.PermissionLock.RLock()
	permission, exists := handlers.PermissionMap[req]
	handlers.PermissionLock.RUnlock()

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

func logRequest(req models.Request) error {
	file, err := os.OpenFile("request.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// format the log entry
	logEntry := fmt.Sprintf("Timestamp: not yet, Method: %s, Path: %s\n", req.Method, req.Path)

	// write the log entry to the file
	_, err = file.WriteString(logEntry)
	if err != nil {
		return err
	}
	return nil
}
