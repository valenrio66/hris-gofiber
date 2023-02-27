package main

import (
	"hris-gofiber/routes"

	"github.com/gofiber/fiber/v2"

	_ "hris-gofiber/docs"

	fiberSwagger "github.com/gofiber/swagger"
)

// @title Human Resource Information System
// @version 1.0
// @description HRIS Go Fiber API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:80
// @BasePath /v1
func main() {
	app := fiber.New()
	app.Get("/swagger/*", fiberSwagger.HandlerDefault)

	routes.SetupRoutes(app)

	app.Listen(":80")
}
