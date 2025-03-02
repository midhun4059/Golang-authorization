package main

import (
	"golang-authorization/config"
	"golang-authorization/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()


	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":3000")
}