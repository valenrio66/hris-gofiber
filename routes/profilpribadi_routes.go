package routes

import (
	"hris-gofiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	setupRoutes := v1.Group("/hris")
	setupRoutes.Post("/getToken", controllers.GetTokenHandler)
}
