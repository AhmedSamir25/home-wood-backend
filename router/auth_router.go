package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func AuthRouter(app *fiber.App) {
	app.Post("/signup", controller.CreateUser)
	app.Get("/login", controller.LoginUser)
	app.Post("/forgetpassword", controller.SendToken)
}
