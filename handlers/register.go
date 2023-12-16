package handlers

import (
	"log"

	"github.com/KirigayaKazuto91/go-login/databases"
	"github.com/KirigayaKazuto91/go-login/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func RegisterPage(c *fiber.Ctx) error {
	return c.SendFile("./templates/registrasi.html")
}

func StoreUser(c *fiber.Ctx) error {
	db, err := databases.ConnectDB()
	if err != nil{
		log.Fatal(err)
	}
	models.MigrateUsers(db)

		username := c.FormValue("username")
		password := c.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil{
			return c.SendStatus(fiber.StatusNotAcceptable)
		}
	
		newUser := User{Username: username, Password: string(hashedPassword)}
		result := db.Create(&newUser)
		if result.Error != nil{
			return c.SendStatus(fiber.StatusExpectationFailed)
		}
		return c.SendStatus(fiber.StatusCreated)
}