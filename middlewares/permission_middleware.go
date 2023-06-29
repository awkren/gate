package middlewares

import (
	"fmt"
	"gate/handlers"
	"gate/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func CheckPermissionMiddleware(c *gin.Context) {
	req := models.Request{
		Method: c.Request.Method,
		Path:   c.Request.URL.Path,
	}

	// Send the request to the permission check channel
	handlers.Requests <- req

	handlers.PermissionLock.RLock()
	permission, exists := handlers.PermissionMap[req]
	handlers.PermissionLock.RUnlock()

	// log the request into a file
	err := logRequest(req, permission)
	if err != nil {
		fmt.Println("Error man kkkkkk")
	}

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

func logRequest(req models.Request, permission bool) error {
	file, err := os.OpenFile("request.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// get timestamp
	getTime := time.Now()

	// determine access status based on permission
	status := "DENIED"
	if permission {
		status = "APPROVED"
	}

	// format the log entry
	logEntry := fmt.Sprintf("Timestamp: %s, Method: %s, Path: %s | Status: %s\n", getTime.Format(time.RFC822), req.Method, req.Path, status)

	// write the log entry to the file
	_, err = file.WriteString(logEntry)
	if err != nil {
		return err
	}
	return nil
}
