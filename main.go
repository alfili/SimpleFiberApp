package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"

	"simplefiberapp/db"
	"simplefiberapp/handlers"
	"simplefiberapp/tools"
)

var appPort string = "5000"

func main() {

	tools.Store = tools.Storage{Sessions: session.New()}

	db.ConnectDB()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// middleware для аутентификации
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println(c.Cookies("Token"))

		sessCookie, _ := tools.Store.Sessions.Storage.Get(c.Cookies("Token"))

		fmt.Println("SESSION: " + string(sessCookie))
		return c.Next()
	})

	SetupRouter(app)

	app.Listen(":" + appPort)
}

func SetupRouter(a *fiber.App) {
	a.Get("/", handlers.HomePage)

	a.Get("/mods", handlers.ModsPage)
	a.Get("/mods/:id", handlers.GetOneModPage)
	a.Get("/create/mod", handlers.ModsAddForm)
	a.Post("/create/mod", handlers.ModsAddPost)

	a.Get("/login", handlers.LoginPage)
	a.Post("/login", handlers.LoginUser)

	a.Get("/reg", handlers.RegisterPage)
	a.Post("/reg", handlers.RegisterUser)
}
