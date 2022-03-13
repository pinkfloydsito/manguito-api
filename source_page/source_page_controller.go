package source_page

import "github.com/gofiber/fiber/v2"

func NewSourcePage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"souce_page": "not implemented",
	})
}

func GetSourcePage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"souce_page": "not implemented",
	})
}

func GetSourcePages(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"souce_page": "not implemented",
	})
}

func UpdateSoucePage(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"souce_page": "not implemented",
	})
}

func DeleteSoucePage(c *fiber.Ctx) error {
	return nil
}
