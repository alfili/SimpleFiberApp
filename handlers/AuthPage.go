package handlers

import (
	"fmt"
	"simplefiberapp/db"
	"simplefiberapp/models"
	"simplefiberapp/tools"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

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

func RegisterUser(c *fiber.Ctx) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	passHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return c.SendString("Не удалось зарегистрировать! Попробуйте снова")
	}

	newUser := models.User{Username: username, Password: string(passHash)}

	db.DBConn.Create(&newUser)

	return c.Render("pages/register", fiber.Map{
		"Title": "Войти",
	}, "layouts/main")
}

func LoginUser(c *fiber.Ctx) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	var user models.User
	db.DBConn.Where("username = ?", username).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return c.SendString("Не удалось войти! Попробуйте снова")
	}

	err = tools.Store.Sessions.Storage.Set("user", []byte(username), time.Hour*12)
	if err != nil {
		return c.SendString("Не удалось войти! Попробуйте снова")
	}

	sessUser, err := tools.Store.Sessions.Storage.Get("user")
	if err != nil {
		return c.SendString("Не удалось войти! Попробуйте снова")
	}

	fmt.Println(string(sessUser))

	return c.SendString("Удалось войти!")
}
