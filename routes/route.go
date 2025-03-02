package routes

import (
	"golang-authorization/controllers"
	"golang-authorization/middleware"

	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/register", controllers.CreateUser)
	router.POST("/login", controllers.LoginUser)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
			protected.GET("/users", controllers.GetUsers)
			protected.PUT("/users/:id", controllers.UpdateUser)
			protected.DELETE("/users/:id", controllers.DeleteUser)
	}
	
}