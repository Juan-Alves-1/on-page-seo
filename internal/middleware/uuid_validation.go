package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UUIDValidationnMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No UUID provided"})
			c.Abort()
			return
		} else {
			uuidToken := strings.TrimPrefix(authHeader, "Bearer ")
			c.Set("session_uuid", uuidToken)
		}

		c.Next()

	}
}
