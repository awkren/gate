package middlewares

import (
	"fmt"
	"net"
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
	// testing
	// if ip == "127.0.0.1" || ip == "::1" {
	// 	return true
	// }

	// check if the ip ip is in the allowedIPs slice
	for _, allowedIP := range allowedIPs {
		if ip == allowedIP {
			return true
		}
	}

	// check if the ip is a valid IPv4 or IPv6 address
	parsedIP := net.ParseIP(ip)
	if parsedIP != nil {
		for _, allowedIP := range allowedIPs {
			if parsedIP.Equal(net.ParseIP(allowedIP)) {
				return true
			}
		}
	}

	return false
}
