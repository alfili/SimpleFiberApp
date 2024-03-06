package middleware

import (
	"simplefiberapp/db"
	"simplefiberapp/models"
	"simplefiberapp/tools"

	"github.com/gofiber/fiber/v2"
)

func Me(c *fiber.Ctx) error {
	cookieId, err := tools.Store.Sessions.Storage.Get(c.Cookies("Token"))
	if err != nil {
		return c.Next()
	}

	if string(cookieId) == "" {
		c.Locals("user", nil)
		return c.Next()
	}

	var user models.User
	db.DBConn.Where("id = ?", cookieId).First(&user)

	c.Locals("user", &user)

	return c.Next()
}
