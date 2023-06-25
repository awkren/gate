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

// original code

// func handleUserRequest(c *gin.Context) {
// 	// forward the request to the external API if access is granted
// 	if permissionGranted() {
// 		http.Redirect(c.Writer, c.Request, externalAPI+endpoint, http.StatusTemporaryRedirect)
// 	} else {
// 		// Respond with an error message or denial
// 		c.JSON(http.StatusForbidden, gin.H{
// 			"error": "Permission denied :p",
// 		})
// 	}
// }
//
