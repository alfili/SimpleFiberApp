package handlers

import "github.com/gofiber/fiber/v2"

func HomePage(c *fiber.Ctx) error {
	return c.Render("pages/index", fiber.Map{
		"Title": "Главная",
	}, "layouts/main")
}
