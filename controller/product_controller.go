package controller

import (
	"homewood/database"
	"homewood/model"

	"github.com/gofiber/fiber/v3"
)

func GetAllProducts(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "get all product be success",
		"statusText": "Ok",
	}
	db := database.DbConn
	var record []model.Products
	db.Find(&record)
	if record == nil {
		context["msg"] = "error when get products"
		context["statusText"] = "error"
		c.Status(400)
		return c.JSON(context)
	}

	context["products"] = record
	c.Status(200)
	return c.JSON(context)
}

func AddProduct(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "success add product",
		"statusText": "Ok",
	}
	record := new(model.Products)
	if err := c.Bind().Body(record); err != nil {
		context["statusText"] = "bad"
		context["msg"] = "invalid request"
		c.Status(400)
		return c.JSON(context)
	}
	db := database.DbConn
	result := db.Create(record)
	if result.Error != nil {
		context["msg"] = "error when add product"
		context["statusText"] = "error"
		c.Status(400)
		c.JSON(context)
	}
	c.Status(200)
	return c.JSON(context)
}
