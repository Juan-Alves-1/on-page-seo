package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UUIDCreationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			newUUID := uuid.New().String()
			c.Header("Authorization", "Bearer "+newUUID)
			c.Set("session_uuid", newUUID)
		} else {
			uuidToken := strings.TrimPrefix(authHeader, "Bearer ")
			c.Set("session_uuid", uuidToken)
		}

		c.Next()

	}
}
