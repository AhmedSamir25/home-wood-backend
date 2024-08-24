package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func CategoriesRouter(app *fiber.App) {
	app.Get("homewood/categories", controller.GetCategories)
	app.Post("homewood/categories", controller.AddCategory)
	app.Delete("homewood/categories/:id", controller.DeleteCategory)
	app.Put("homewood/categories/:id", controller.UpdateCategory)
}
