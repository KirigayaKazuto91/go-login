package handlers

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func RegisterPage(w http.ResponseWriter, r *http.Request){
	dsn := "host=localhost user=postgres password=root dbname=projectone port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})
	
	
	if r.Method != http.MethodPost{
		http.ServeFile(w, r, "./templates/registrasi.html")
		return
	}

		username := r.FormValue("username")
		password := r.FormValue("password")

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil{
			http.Error(w, "Failed to hash the password", http.StatusInternalServerError)
			return
		}
	
		newUser := User{Username: username, Password: string(hashedPassword)}
		result := db.Create(&newUser)
		if result.Error != nil{
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			return
		}
		
		fmt.Fprintf(w, "User Registered successfully!")
		return
}