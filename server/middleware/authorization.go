package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang-authorization/models"
	"golang-authorization/config"


)

// Check if user has the required permission
func PermissionMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing or invalid"})
			c.Abort()
			return
		}

		userClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format from permissionMiddleware"})
			c.Abort()
			return
		}

		userID := uint(userClaims["user_id"].(float64))

		// Fetch user from DB along with their role and permissions
		var user models.User
		if err := config.DB.Preload("Role.Permissions").First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Check if user has the required permission
		for _, permission := range user.Role.Permissions {
			if permission.Permissions == requiredPermission {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Not Authorized"})
		c.Abort()
	}
}