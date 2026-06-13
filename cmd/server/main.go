package main

import (
	"log"

	"user-api/database"
	"user-api/internal/logger"
	"user-api/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	logger.InitLogger()

	app := fiber.New()

	routes.SetupRoutes(app)

	logger.Log.Info("Server starting on port 8080")

	log.Fatal(app.Listen(":8080"))

}
