package handlers

import (
	"log"
	"net/http"

	"github.com/KirigayaKazuto91/go-login/databases"
	"github.com/KirigayaKazuto91/go-login/models"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
  db, err := databases.ConnectDB()
  if err != nil {
    log.Fatal(err)
  }
  models.MigrateUsers(db)
  
  if r.Method != http.MethodPost {
	http.ServeFile(w, r, "./templates/login.html")
	return
  }

  username := r.FormValue("username")
  password := r.FormValue("password")

  var user models.User

  if err := db.Where("username = ?", username).First(&user).Error; err != nil{
    http.Error(w, "Invalid Username or password", http.StatusUnauthorized)
    return
  }

  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
    http.Error(w, "Invalid Username or password", http.StatusUnauthorized)
    return
  }

	http.Redirect(w, r, "/home", http.StatusFound)

}
