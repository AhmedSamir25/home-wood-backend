package controller

import (
	"homewood/database"
	"homewood/model"

	"github.com/gofiber/fiber/v3"
)

func AddOrder(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "order successfully",
		"statusText": "Ok",
	}
	record := new(model.Orders)
	db := database.DbConn
	if err := c.Bind().Body(record); err != nil {
		context["msg"] = "invalid request"
		context["statusText"] = "error"
		c.Status(400)
		return c.JSON(context)
	}
	result := db.Create(record)
	if result.Error != nil {
		context["msg"] = "An error occurred while placing an order"
		context["statusText"] = "error"
		return c.Status(400).JSON(context)
	}

	return c.Status(200).JSON(context)
}
