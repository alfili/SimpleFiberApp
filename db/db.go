package db

import (
	"fmt"
	"simplefiberapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ConnectDB() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/simplefiber?charset=utf8mb4&parseTime=True&loc=Local"

	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Не удалось подключиться к базе данных: " + err.Error())
	}

	fmt.Println("Подключились к базе данных!")

	DBConn.AutoMigrate(&models.Mod{})
	DBConn.AutoMigrate(&models.User{})
	DBConn.AutoMigrate(&models.Profile{})

	fmt.Println("База данных мигрирована!")
}
