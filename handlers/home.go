package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HomePage(c *fiber.Ctx) error {
	return c.SendFile("./templates/home.html")
}