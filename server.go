package main

import (
	"homewood/database"
	"homewood/router"

	"github.com/gofiber/fiber/v3"
)

func main() {
	database.ConnDB()
	app := fiber.New()
	router.AuthRouter(app)
	app.Listen(":3002")
}
