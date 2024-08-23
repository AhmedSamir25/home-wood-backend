package controller

import (
	"homewood/database"
	"homewood/model"
	"log"

	"github.com/gofiber/fiber/v3"
)

func AddBanner(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "Successful banner added",
		"statusText": "Ok",
	}
	credentials := new(model.Banner)
	if err := c.Bind().Body(credentials); err != nil {
		log.Printf("Error parsing request body: %v", err)
		context["statusText"] = "bad"
		context["msg"] = "invalid request"
		c.Status(400)
		return c.JSON(context)
	}
	if credentials.BannerImage == "" {
		context["statusText"] = "bad"
		context["msg"] = "Banner Image is required"
		c.Status(400)
		return c.JSON(context)
	}
	result := database.DbConn.Create(credentials)
	if result.Error != nil {
		log.Println("Error in saving data:", result.Error)
		context["statusText"] = "bad"
		context["msg"] = "error in save"
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(context)
	}
	c.Status(200)
	return c.JSON(context)
}

func GetBanners(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "succees",
		"statusText": "Ok",
	}
	db := database.DbConn
	var records []model.Banner
	db.Find(&records)
	context["banners"] = records
	c.Status(200)
	return c.JSON(context)
}

func DeleteBanner(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "delete banner success",
		"statusText": "Ok",
	}
	id := c.Params("id")
	record := new(model.Banner)
	result := database.DbConn.First(&record, "banner_id = ?", id)

	if result.Error != nil {
		context["msg"] = "record not found"
		context["statusText"] = "error"
		return c.Status(fiber.StatusNotFound).JSON(context)
	}

	result = database.DbConn.Where("banner_id = ?", id).Delete(&model.Banner{})
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
