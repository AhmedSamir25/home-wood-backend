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

func DeleteCategory(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "delete category success",
		"statusText": "Ok",
	}
	id := c.Params("id")
	record := new(model.Categories)
	result := database.DbConn.First(&record, "category_id = ?", id)

	if result.Error != nil {
		context["msg"] = "record not found"
		context["statusText"] = "error"
		return c.Status(fiber.StatusNotFound).JSON(context)
	}

	result = database.DbConn.Where("category_id = ?", id).Delete(&model.Categories{})
	if result.Error != nil {
		context["msg"] = "error deleting record"
		context["statusText"] = "error"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	if result.RowsAffected == 0 {
		context["msg"] = "no records were deleted"
		context["statusText"] = "warning"
		return c.Status(fiber.StatusNotFound).JSON(context)
	}

	return c.Status(fiber.StatusOK).JSON(context)
}

func UpdateCategory(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "update category success",
		"statusText": "Ok",
	}
	id := c.Params("id")
	record := new(model.Categories)

	result := database.DbConn.First(record, "category_id = ?", id)
	if result.Error != nil {
		context["msg"] = "record not found"
		context["statusText"] = "error"
		return c.Status(fiber.StatusNotFound).JSON(context)
	}

	if err := c.Bind().Body(record); err != nil {
		context["msg"] = "error parsing request body"
		context["statusText"] = "error"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}

	result = database.DbConn.Model(record).Where("category_id = ?", id).Save(record)
	if result.Error != nil {
		context["msg"] = "error updating record"
		context["statusText"] = "error"
		return c.Status(fiber.StatusInternalServerError).JSON(context)
	}

	if result.RowsAffected == 0 {
		context["msg"] = "no records were updated"
		context["statusText"] = "warning"
		return c.Status(fiber.StatusNotFound).JSON(context)
	}
	context["category"] = record
	return c.Status(fiber.StatusOK).JSON(context)
}
