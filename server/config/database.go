
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"golang-authorization/models" 
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DB_URL")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = database
	fmt.Println("Database connected successfully!")

	// Correcting the variable name
	DB.AutoMigrate(&models.Role{}, &models.Permission{}, &models.RolePermission{}, &models.User{})
	fmt.Println("Database Migrated Successfully!")
}