package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Check if user has the required permission
func PermissionMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userClaims := claims.(*jwt.MapClaims)
		userPermissions, ok := (*userClaims)["permissions"].([]any)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid permissions"})
			c.Abort()
			return
		}

		for _, perm := range userPermissions {
			if perm.(string) == requiredPermission {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Not Authorized"})
		c.Abort()
	}
}