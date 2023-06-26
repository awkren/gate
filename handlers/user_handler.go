package handlers

import (
	"gate/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var (
	PermissionMap  map[models.Request]bool
	PermissionLock sync.RWMutex
	Requests       chan models.Request
	ExternalAPI    string
	Endpoint       string
	UserEndpoint   string
)

func HandleUserRequest(c *gin.Context) {
	PermissionGranted := func() bool {
		PermissionLock.RLock()
		permission := PermissionMap[models.Request{Method: "GET", Path: "/users"}]
		PermissionLock.RUnlock()
		return permission
	}
	// forward the request
	if PermissionGranted() {
		http.Redirect(
			c.Writer,
			c.Request,
			ExternalAPI+Endpoint,
			http.StatusTemporaryRedirect,
		)
	} else {
		// respond with an error message or denial
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Permission denied :p",
		})
	}
}
