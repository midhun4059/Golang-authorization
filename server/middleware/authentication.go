package middleware

import (
	"golang-authorization/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token msg from authMiddleware"})
			c.Abort()
			return
		}

		// Extract Bearer token
		parts := strings.Split(tokenHeader, " ")
fmt.Println("parts from authMiddleware:", parts)

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format msg from authMiddleware"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		fmt.Println("Token received:", tokenString) // Debugging

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token from auth middleware "})
				c.Abort()
				return
		}
		c.Set("user", claims)
		c.Next()
	}
}