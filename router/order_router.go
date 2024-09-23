package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func OrderRouter(app *fiber.App) {
	app.Post("homewood/order", controller.AddOrder)
}
