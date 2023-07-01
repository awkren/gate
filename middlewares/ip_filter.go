package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func IPFilterMiddleware(c *gin.Context) {
	allowedIPs := strings.Split(os.Getenv("ALLOWED_IPS"), ",")

	clientIP := c.ClientIP()
	fmt.Println(clientIP)

	// check if the client's is in the allowedIPs slice
	if !isIPAllowed(clientIP, allowedIPs) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied :p",
		})
		c.Abort()
		return
	}

	// if the ip is allowed, continue processing the request
	c.Next()
}

func isIPAllowed(ip string, allowedIPs []string) bool {
	// check if the ip ip is in the allowedIPs slice
	for _, allowedIP := range allowedIPs {
		if ip == allowedIP {
			return true
		}
	}
	return false
}
