package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"simplefiberapp/db"
	"simplefiberapp/handlers"
)

var appPort string = "5000"

func main() {

	db.ConnectDB()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	SetupRouter(app)

	app.Listen(":" + appPort)
}

func SetupRouter(a *fiber.App) {
	a.Get("/", handlers.HomePage)

	a.Get("/mods", handlers.ModsPage)
	a.Get("/mods/:id", handlers.GetOneModPage)
	//a.Get("/mods/add", handlers.ModsAddForm)
	a.Post("/mods", handlers.ModsAddPost)
}
