package handlers

import (
	"fmt"
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
		return c.SendString("Не удалось зарегистрировать! Попробуйте снова")
	}

	newUser := models.User{Username: username, Password: string(passHash)}

	db.DBConn.Create(&newUser)

	return c.Render("pages/register", fiber.Map{
		"Title": "Войти",
	}, "layouts/main")
}

// Сессии можно хранить в приложении в формате "токен куки": "id"
func LoginUser(c *fiber.Ctx) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	var user models.User
	db.DBConn.Where("username = ?", username).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return c.SendString("Не удалось войти! Попробуйте снова")
	}

	c.Cookie(&fiber.Cookie{
		Name:    "Token",
		Value:   uuid.NewString(),
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 12),
	})

	err = tools.Store.Sessions.Storage.Set(c.Cookies("Token"), []byte(strconv.Itoa(int(user.ID))), time.Hour*12)
	if err != nil {
		return c.SendString("Не удалось войти! Попробуйте снова")
	}

	return c.SendString("Удалось войти!")
}

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

	fmt.Println(c.Locals("user"))

	return c.Next()
}
