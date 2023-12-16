package handlers

import (
	"log"

	"github.com/KirigayaKazuto91/go-login/databases"
	"github.com/KirigayaKazuto91/go-login/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("123456"))

func LoginPage(c *fiber.Ctx) error {
  return c.SendFile("./templates/login.html")
}


func CheckLogin(c *fiber.Ctx) error {
  
  db, err := databases.ConnectDB()
  if err != nil {
    log.Fatal(err)
  }
  models.MigrateUsers(db)
  
    username := c.FormValue("username")
    password := c.FormValue("password")
  
    var user models.User
  
    if err := db.Where("username = ?", username).First(&user).Error; err != nil{
      return c.SendStatus(fiber.StatusUnauthorized)
    }
  
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
      return c.SendStatus(fiber.StatusUnauthorized)
    }
  
    return c.Redirect("/home")
}
