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
	router.CategoriesRouter(app)
	router.ProductsRouting(app)
	router.FavoriteRouting(app)
	app.Listen("0.0.0.0:3002")
}
