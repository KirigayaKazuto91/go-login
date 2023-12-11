package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KirigayaKazuto91/go-login/handlers"
)

func main() {
	http.HandleFunc("/login", handlers.LoginPage)
	http.HandleFunc("/home", handlers.HomePage)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js"))))

	fmt.Printf("Starting server at port 8080\n")
	
	err := http.ListenAndServe(":8080", nil)
	if  err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}