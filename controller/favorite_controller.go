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
