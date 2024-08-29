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

func UpdateProduct(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "update product success",
		"statusText": "Ok",
	}
	id := c.Params("id")
	record := new(model.Products)
	result := database.DbConn.First(record, "product_id = ?", id)
	if result.Error != nil {
		context["msg"] = "record not found"
		context["statusText"] = "error"
		c.Status(400)
	}
	if err := c.Bind().Body(record); err != nil {
		context["msg"] = "error parsing request body"
		context["statusText"] = "error"
		c.Status(400)
	}
	result = database.DbConn.Model(record).Where("product_id = ?", id).Save(record)
	if result.Error != nil {
		context["msg"] = "error when product"
		context["statusText"] = "error"
		c.Status(400)
	}
	if result.RowsAffected == 0 {
		context["msg"] = "no records were updated"
		context["statusText"] = "warning"
		return c.Status(fiber.StatusNotFound).JSON(context)
	}
	c.Status(200)
	return c.JSON(context)
}

func DeleteProduct(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "delete product success",
		"statusText": "Ok",
	}
	id := c.Params("id")
	record := new(model.Products)
	result := database.DbConn.First(&record, "product_id = ?", id)

	if result.Error != nil {
		context["msg"] = "record not found"
		context["statusText"] = "error"
		return c.Status(fiber.StatusNotFound).JSON(context)
	}
	result = database.DbConn.Where("product_id = ?", id).Delete(&model.Products{})
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
	c.Status(200)
	return c.JSON(context)
}
