package handlers

import (
	"simplefiberapp/db"
	"simplefiberapp/models"
	"simplefiberapp/tools"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		return c.Render("pages/register", fiber.Map{
			"Title":  "Ругистрация",
			"Errors": []string{"Не удалось зарегистрироваться! Попробуйте снова"},
		}, "layouts/main")
	}

	newUser := models.User{Username: username, Password: string(passHash)}

	db.DBConn.Create(&newUser)

	return c.Redirect("/login")
}

// Сессии можно хранить в приложении в формате "токен куки": "id"
func LoginUser(c *fiber.Ctx) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	var user models.User
	db.DBConn.Where("username = ?", username).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return c.Render("pages/login", fiber.Map{
			"Title":  "Войти",
			"Errors": []string{"Не удалось войти! Попробуйте снова"},
		}, "layouts/main")
	}

	c.Cookie(&fiber.Cookie{
		Name:    "Token",
		Value:   uuid.NewString(),
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 12),
	})

	err = tools.Store.Sessions.Storage.Set(c.Cookies("Token"), []byte(strconv.Itoa(int(user.ID))), time.Hour*12)
	if err != nil {
		return c.Render("pages/login", fiber.Map{
			"Title":  "Войти",
			"Errors": []string{"Не удалось войти! Попробуйте снова"},
		}, "layouts/main")
	}

	return c.Redirect("/")
}
