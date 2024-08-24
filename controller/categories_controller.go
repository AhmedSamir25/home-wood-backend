package controller

import (
	"homewood/database"
	"homewood/model"

	"github.com/gofiber/fiber/v3"
)

func GetCategories(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "get all categories succees",
		"statusText": "Ok",
	}

	db := database.DbConn
	var record []model.Categories
	db.Find(&record)
	if record == nil {
		context["msg"] = "error when get categories"
		context["statusText"] = "error"
		c.Status(400)
		return c.JSON(context)
	}
	context["categories"] = record
	c.Status(200)
	return c.JSON(context)
}

func AddCategory(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "add category succees",
		"statusText": "Ok",
	}
	return c.JSON(context)
}
