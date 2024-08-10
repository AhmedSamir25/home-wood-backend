package main

import (
	"homewood/database"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.ConnDB()
	app := fiber.New()
	app.Listen(":3002")
}
