package handlers

import (
	"simplefiberapp/db"
	"simplefiberapp/models"

	"github.com/gofiber/fiber/v2"
)

func UserProfilePage(c *fiber.Ctx) error {

	username := c.Params("username")

	var user models.User
	err := db.DBConn.Preload("UserProfile").Where("username = ?", username).First(&user).Error
	if err != nil {
		return c.Redirect("/")
	}

	return c.Render("pages/profile", fiber.Map{
		"Title": "Профиль пользователя " + username,
		"User":  user,
	}, "layouts/main")
}
