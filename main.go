package main

import (
	"fmt"
	"manguito/v0/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	SetupRoutes(app)
	Serve(app)
}

func SetupRoutes(app *fiber.App) {
	routes.BuildRoutes(app)
}

func Serve(app *fiber.App) {
	if err := app.Listen(":3000"); err != nil {
		fmt.Errorf("Can't listen: %v", err)
	}
}