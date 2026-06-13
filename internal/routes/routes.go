package routes

import (
	handler "user-api/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Post("/users", handler.CreateUser)
	app.Get("/users", handler.GetAllUsers)
	app.Get("/users/:id", handler.GetUserByID)
	app.Put("/users/:id", handler.UpdateUser)
	app.Delete("/users/:id", handler.DeleteUser)
}
