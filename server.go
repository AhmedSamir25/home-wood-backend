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
	router.BannerRouter(app)
	app.Listen("0.0.0.0:3002")
}
