package router

import (
	"homewood/controller"

	"github.com/gofiber/fiber/v3"
)

func AuthRouter(app *fiber.App) {
	app.Post("/homewood/signup", controller.CreateUser)
	app.Get("/homewood/login", controller.LoginUser)
	app.Post("/homewood/forgetpassword", controller.SendToken)
	app.Get("/homewood/checktoken", controller.CheckToken)
	app.Put("/homewood/resetpassword", controller.ResetAndUpdatePassword)
	app.Get("/home/test", controller.TestT)
}
