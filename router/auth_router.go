package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func AuthRouter(app *fiber.App) {
	app.Post("/homewood/signup", controller.CreateUser)
	app.Get("/homewood/login", controller.LoginUser)
	app.Post("/homewood/forgetpassword", controller.SendToken)
	app.Get("/homewood/sendtoken", controller.ResetPassword)
	app.Put("/homewood/resetpassword", controller.UpdatePassword)
	app.Get("/home/test", controller.TestT)
}
