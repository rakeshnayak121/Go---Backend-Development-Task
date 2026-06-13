package main

import (
	"log"
	"user-api/database"
	"user-api/logger"
	"user-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	logger.InitLogger()

	router := gin.Default()

	routes.SetupRoutes(router)
	logger.Log.Info("Server starting on port 8080")

	router.Run(":8080")
}
