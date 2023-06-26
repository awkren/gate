package middlewares

import (
	"gate/handlers"
	"gate/models"
	"github.com/gin-gonic/gin"
	"net/http"
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
