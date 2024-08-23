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
