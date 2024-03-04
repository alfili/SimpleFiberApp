package handlers

import (
	"simplefiberapp/db"
	"simplefiberapp/models"

	"github.com/gofiber/fiber/v2"
)

func ModsPage(c *fiber.Ctx) error {

	db := db.DBConn

	var modList []models.Mod
	db.Find(&modList)

	return c.Render("pages/mods", fiber.Map{
		"Title": "Моды",
		"Mods":  modList,
	}, "layouts/main")
}

func GetOneModPage(c *fiber.Ctx) error {

	id := c.Params("id")

	db := db.DBConn

	var mod models.Mod
	db.Find(&mod, id)

	return c.Render("pages/one_mod", fiber.Map{
		"Title": mod.Name,
		"Mod":   mod,
	}, "layouts/main")
}

func ModsAddPost(c *fiber.Ctx) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	db := db.DBConn

	var mod models.Mod = models.Mod{Name: name, Description: description}
	db.Create(&mod)

	return c.Redirect("/mods/")
}
