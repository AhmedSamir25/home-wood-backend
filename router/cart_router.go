package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func CartRouter(app *fiber.App) {
	app.Post("homewood/cart", controller.AddProductToCart)
}
