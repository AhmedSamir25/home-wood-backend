package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func BannerRouter(app *fiber.App) {
	app.Post("homewood/banner", controller.AddBanner)
}
