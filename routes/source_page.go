package routes

import (
	"github.com/gofiber/fiber/v2"

	"manguito/v0/source_page"
)

func SourcePageRoutes(route fiber.Router) {
	route.Get("/source_pages", source_page.GetSourcePages)
	route.Get("/source_pages/:id", source_page.GetSourcePage)

	route.Post("/source_pages/new", source_page.NewSourcePage)
	route.Delete("/source_pages/", source_page.DeleteSoucePage)
}


func BuildRoutes(app *fiber.App) {
	route := app.Group("api/v1")

	SourcePageRoutes(route)
}
