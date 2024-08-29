package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func ProductsRouting(app *fiber.App) {
	app.Get("homewood/products", controller.GetAllProducts)
	app.Post("homewood/product", controller.AddProduct)
	app.Put("homewood/product/:id", controller.UpdateProduct)
	app.Delete("homewood/product/:id", controller.DeleteProduct)
}
