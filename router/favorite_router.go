package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func FavoriteRouting(app *fiber.App) {
	app.Get("homewood/favorite/products/user=:id", controller.GetFavoriteProducts)
}
