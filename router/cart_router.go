package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func CartRouter(app *fiber.App) {
	app.Post("homewood/cart", controller.AddProductToCart)
	app.Get("homewood/cart/user=:id", controller.GetProductFromCart)
	app.Put("homewood/cart/product=:id", controller.UpdateProductQt)
}
