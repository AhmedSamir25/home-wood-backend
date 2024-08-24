package controller

import (
	"homewood/database"
	"homewood/model"
	"log"

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
	record := new(model.Categories)
	if err := c.Bind().Body(record); err != nil {
		log.Printf("Error parsing request body: %v", err)
		context["statusText"] = "bad"
		context["msg"] = "invalid request"
		c.Status(400)
		return c.JSON(context)
	}
	if record.CategoryName == "" {

		context["msg"] = "name is empty"
		context["statusText"] = "error"
		c.Status(400)
		return c.JSON(context)
	}
	result := database.DbConn.Create(record)
	if result.Error != nil {
		context["msg"] = "error when add categories"
		context["statusText"] = "error"
		c.Status(400)
		c.JSON(context)
	}
	c.Status(200)
	return c.JSON(context)
}
