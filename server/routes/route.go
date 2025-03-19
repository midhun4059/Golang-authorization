package routes

import (
	"golang-authorization/controllers"
	"golang-authorization/middleware"

	
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", controllers.LoginUser)
	router.POST("/refresh",controllers.RefreshToken)
	
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		  protected.POST("/register",middleware.PermissionMiddleware("Create"), controllers.CreateUser)
			protected.GET("/users",middleware.PermissionMiddleware("Read"), controllers.GetUsers)
			protected.PUT("/users/:id",middleware.PermissionMiddleware("Update"), controllers.UpdateUser)
			protected.DELETE("/users/:id",middleware.PermissionMiddleware("Delete"), controllers.DeleteUser)
	}
	
}

//here middleware.AuthMiddleware is used because it wants to call in all routes but in the case of permissionMiddle ware we have to give specific permission to specofic routes.