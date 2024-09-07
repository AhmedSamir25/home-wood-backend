package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func ProductsRouting(app *fiber.App) {
	app.Get("homewood/products/:pageid", controller.GetAllProducts)
	app.Post("homewood/product", controller.AddProduct)
	app.Put("homewood/product/:id", controller.UpdateProduct)
	app.Delete("homewood/product/:id", controller.DeleteProduct)
	app.Get("homewood/products/category/:id", controller.GetProductsPyCategories)
	app.Get("homewood/product/details/:id", controller.GetProductDetails)
}
