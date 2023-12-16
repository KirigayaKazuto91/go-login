package main

import (
	"net/http"

	"github.com/KirigayaKazuto91/go-login/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/", "./templates")
	
	app.Get("/home", handlers.HomePage)
	app.Get("/login", handlers.LoginPage)
	app.Post("/login", handlers.CheckLogin)

	app.Get("/register", handlers.RegisterPage)
	app.Post("/register", handlers.StoreUser)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js"))))

	app.Listen(":8080")

}