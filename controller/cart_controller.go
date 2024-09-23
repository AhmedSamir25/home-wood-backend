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
