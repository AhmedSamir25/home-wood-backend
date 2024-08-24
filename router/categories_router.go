package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func CategoriesRouter(app *fiber.App) {
	app.Get("homewood/categories", controller.GetCategories)
	app.Post("homewood/categories", controller.AddCategory)
}
