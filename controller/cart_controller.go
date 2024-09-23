package controller

import (
	"homewood/database"
	"homewood/model"

	"github.com/gofiber/fiber/v3"
)

func AddProductToCart(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "The product has been successfully added to your cart",
		"statusText": "Ok",
	}
	record := new(model.Cart)
	if err := c.Bind().Body(record); err != nil {
		context["msg"] = "invalid request"
		context["statusText"] = "bad"
		c.Status(400)
		return c.JSON(context)
	}
	db := database.DbConn
	result := db.Where("product_id = ? AND user_id = ?", record.ProductId, record.UserId).First(&model.Cart{})
	if result.Error == nil {
		if err := db.Where("product_id = ? AND user_id = ?", record.ProductId, record.UserId).Delete(&model.Cart{}).Error; err != nil {
			context["msg"] = "Error deleting the product from cart"
			context["statusText"] = "error"
			return c.Status(fiber.StatusInternalServerError).JSON(context)
		}

		context["msg"] = "Product has been successfully removed from cart"
		context["statusText"] = "Ok"
		return c.Status(fiber.StatusOK).JSON(context)
	}

	if err := db.Create(record).Error; err != nil {
		context["msg"] = "An error occurred while adding the product to cart"
		context["statusText"] = "bad"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}
	return c.JSON(context)
}

func GetProductFromCart(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "All products have been successfully fetched from the cart",
		"statusText": "Ok",
	}
	userId := c.Params("id")
	db := database.DbConn
	var record []model.Products
	err := db.Table("products").
		Select("*").
		Joins("JOIN carts AS f ON f.product_id = products.product_id").
		Where("f.user_id = ?", userId).
		Find(&record).Error
	if err != nil {
		context["msg"] = "error when get products"
		context["statusText"] = "error"
		c.Status(400)
		return c.JSON(context)
	}
	context["products"] = record
	return c.JSON(context)
}
