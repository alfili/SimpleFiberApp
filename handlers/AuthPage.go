package handlers

import "github.com/gofiber/fiber/v2"

func LoginPage(c *fiber.Ctx) error {
	return c.Render("pages/login", fiber.Map{
		"Title": "Войти",
	}, "layouts/main")
}

func RegisterPage(c *fiber.Ctx) error {
	return c.Render("pages/register", fiber.Map{
		"Title": "Регистрация",
	}, "layouts/main")
}
