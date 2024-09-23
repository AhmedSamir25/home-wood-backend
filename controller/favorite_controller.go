package controller

import (
	"homewood/database"
	"homewood/model"

	"github.com/gofiber/fiber/v3"
)

func GetFavoriteProducts(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "get all favorite product done",
		"statusText": "Ok",
	}

	userID := c.Params("id")
	db := database.DbConn
	var products []model.Products

	err := db.Table("products").
		Select("*").
		Joins("JOIN favorites AS f ON f.product_id = products.product_id").
		Where("f.user_id = ?", userID).
		Find(&products).Error

	if err != nil {
		context["msg"] = "error when get products"
		context["statusText"] = "error"
		c.Status(400)
		return c.JSON(context)
	}
	context["products"] = products
	return c.JSON(context)
}

func AddProductToFavorite(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "The product has been successfully added to your favorites",
		"statusText": "Ok",
	}
	record := new(model.Favorites)
	if err := c.Bind().Body(record); err != nil {
		context["msg"] = "invalid request"
		context["statusText"] = "bad"
		c.Status(400)
		return c.JSON(context)
	}
	db := database.DbConn

	result := db.Where("product_id = ? AND user_id = ?", record.ProductId, record.UserId).First(&model.Favorites{})

	if result.Error == nil {
		if err := db.Where("product_id = ? AND user_id = ?", record.ProductId, record.UserId).Delete(&model.Favorites{}).Error; err != nil {
			context["msg"] = "Error deleting the product from favorites"
			context["statusText"] = "error"
			return c.Status(fiber.StatusInternalServerError).JSON(context)
		}

		context["msg"] = "Product has been successfully removed from favorites"
		context["statusText"] = "Ok"
		return c.Status(fiber.StatusOK).JSON(context)
	}

	if err := db.Create(record).Error; err != nil {
		context["msg"] = "An error occurred while adding the product to favorites"
		context["statusText"] = "bad"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	}
	return c.JSON(context)
}
