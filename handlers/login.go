package handlers

import (
	"fmt"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
	http.ServeFile(w, r, "./templates/login.html")
	return
  }

  username := r.FormValue("username")
  password := r.FormValue("password")

  if username == "admin" && password == "password"{
	http.Redirect(w, r, "/home", http.StatusFound)
	return
  }

  fmt.Fprintf(w, "Invalid credentials. Please try Again.")
}
