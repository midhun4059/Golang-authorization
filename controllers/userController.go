package controllers

import (
	"golang-authorization/config"
	"golang-authorization/models"
	
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	// Attempt to fetch users from the database
	if err := config.DB.Preload("Role").Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users", "details": err.Error()})
		return
	}

	// If no users are found, return a message instead of an empty array
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users found"})
		return
	}

	// Successfully retrieved users
	c.JSON(http.StatusOK, users)
}

func UpdateUser(c *gin.Context) {
	// Extract user claims from context
	claims, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing or invalid"})
		c.Abort()
		return
	}

	// Correct type assertion
	userClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}

	// Get user role from token
	userRole, ok := userClaims["role"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role in token"})
		c.Abort()
		return
	}

	// Only allow "Super Admin" and "Admin" to update users
	if userRole != "SuperAdmin" && userRole != "Admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not Authorized"})
		c.Abort()
		return
	}

	var user models.User
	id := c.Param("id")

	// Find user by ID
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Parse only provided fields from request body
	var updatedData map[string]interface{}
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Prevent updating sensitive fields like ID
	delete(updatedData, "id")

	// Perform the actual update in the database
	if err := config.DB.Model(&user).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
		return
	}

	// Send success response after updating the database
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}


func DeleteUser(c *gin.Context) {
	// Extract user claims from context
	claims, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing or invalid"})
		c.Abort()
		return
	}

	// Correct type assertion
	userClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}

	// Get user role from token
	userRole, ok := userClaims["role"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role in token"})
		c.Abort()
		return
	}

	// Only allow "SuperAdmin" to delete users
	if userRole != "SuperAdmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not Authorized. Only SuperAdmin can delete users."})
		c.Abort()
		return
	}

	var user models.User
	id := c.Param("id")

	// Check if user exists before deleting
	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	// Delete the user
	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


func CreateUser(c *gin.Context) {
	// Extract user claims from context
	claims, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token missing or invalid msg from update func"})
		c.Abort()
		return
	}

	// Correct type assertion
	userClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}

	// Get user role from token
	userRole, ok := userClaims["role"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role in token"})
		c.Abort()
		return
	}

	// Only allow "SuperAdmin" to create users
	if userRole != "SuperAdmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not Authorized. Only SuperAdmin can create users."})
		c.Abort()
		return
	}

	var user models.User

	// Bind JSON request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Create user in the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	// Send success response
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
