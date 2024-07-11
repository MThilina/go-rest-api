package main

import (
	"go/go-crud/config"
	"go/go-crud/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file")
		}
	}

}

func main() {

	appHost := os.Getenv("APP_HOST")
	appPort := os.Getenv("APP_PORT")

	router := gin.Default()
	config.ConfigureDatabase()
	routes.FetchUser(router) // direct the application to root path
	router.Run(appHost + ":" + appPort)

}
