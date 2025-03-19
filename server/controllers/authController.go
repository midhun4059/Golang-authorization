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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Find user by email and preload Role & Permissions
	var user models.User
	if err := config.DB.Preload("Role.Permissions").Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email ID"})
		return
	}

	// Validate password (direct comparison as per your request)
	if user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Extract permissions from user's role
	var permissions []string
	for _, perm := range user.Role.Permissions {
		permissions = append(permissions, perm.Permissions)
	}

	// Generate JWT Tokens
	accessToken, refreshToken, err := utils.GenerateToken(user.ID, user.Role.EmployeePosition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Save Refresh Token in DB
	user.RefreshToken = refreshToken
	

	// Set Refresh Token as HttpOnly Cookie (secure, not accessible by JavaScript)
	c.SetCookie("refresh_token", refreshToken, 7*24*60*60, "/", "", false, true)

	// Send JSON response to frontend with role & permissions
	c.JSON(http.StatusOK, gin.H{
		"username":     user.Username,
		"access_token": accessToken, 
		"job_title":    user.Role.EmployeePosition,
		"permissions":  permissions, // Send permissions list
	})
}


func RefreshToken(c*gin.Context){
	refreshToken,err :=c.Cookie("refresh_token")

	if err !=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"token is not found in cookies"})
	}


  claims,err :=utils.ValidateRefreshToken(refreshToken)

 if err !=nil{
	c.JSON(http.StatusUnauthorized,gin.H{"error" : "Invalid refresh token to Validate"})
 }

  userID,ok := claims["user_id"].(float64)

 if ! ok{
	c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid token claims from refresh token"})
	return
 }

 var user models.User

 if err :=config.DB.Preload("Role").Where("id= ?",uint(userID)).First(&user).Error;

	err !=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"User not found"})
		return
	}

	
	acesssToken, newRefreshToken ,err :=utils.GenerateToken(user.ID,user.Role.EmployeePosition)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Could  not generate a new token"})
	}

	user.RefreshToken=newRefreshToken
	config.DB.Save(&user)

	c.SetCookie("refresh_token",newRefreshToken,7*24*60*60,"/","",false ,true)

	c.JSON(http.StatusOK,gin.H {"new Access Token": acesssToken})

}

