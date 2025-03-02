package controllers

import (
	"golang-authorization/config"
	"golang-authorization/models"
	"golang-authorization/utils"
	

	"net/http"

	"github.com/gin-gonic/gin"
)



func LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	if err := config.DB.Preload("Role").Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare password directly (Note: Hash passwords in production)
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT with role
	token, err := utils.GenerateToken(user.ID, user.Role.EmployeePosition) // Pass role here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}


// func getRoleName(roleID uint) string {
// 	var role models.Role
// 	if err := config.DB.Where("roll_no = ?", roleID).First(&role).Error; err != nil {
// 		return "Unknown"
// 	}
// 	return role.EmployeePosition
// }