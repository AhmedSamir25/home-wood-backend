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
	app.Listen("192.168.1.5:3002")
}
