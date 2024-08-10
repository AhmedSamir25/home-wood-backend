package controller

import (
	"homewood/database"
	"homewood/model"

	"log"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Create User"}

	record := new(model.User)

	if err := c.Bind().Body(record); err != nil {
		log.Printf("Error parsing request body: %v", err)
		context["statusText"] = "bad"
		context["msg"] = "invalid request"
		c.Status(400)
		return c.JSON(context)
	}

	if record.Name == "" || record.Password == "" {
		context["statusText"] = "bad"
		context["msg"] = "Name and Password are required"
		c.Status(400)
		return c.JSON(context)
	}

	var existingUser model.User
	result := database.DbConn.Where("email = ?", record.Email).First(&existingUser)
	if result.Error == nil {
		context["statusText"] = "bad"
		context["msg"] = "Email already exists"
		c.Status(400)
		return c.JSON(context)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(record.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		context["statusText"] = "bad"
		context["msg"] = "error hashing password"
		c.Status(500)
		return c.JSON(context)
	}
	record.Password = string(hashedPassword)

	result = database.DbConn.Create(record)
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

func LoginUser(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Login successful",
	}

	credentials := new(model.User)

	if err := c.Bind().Body(credentials); err != nil {
		log.Printf("Error parsing request body: %v", err)
		context["statusText"] = "bad"
		context["msg"] = "invalid request"
		c.Status(400)
		return c.JSON(context)
	}

	var user model.User
	result := database.DbConn.Where("email = ?", credentials.Email).First(&user)
	if result.Error != nil {
		context["statusText"] = "bad"
		context["msg"] = "email not found"
		c.Status(400)
		return c.JSON(context)
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		context["statusText"] = "bad"
		context["msg"] = "incorrect password"
		c.Status(400)
		return c.JSON(context)
	}

	context["msg"] = "Login successful"
	c.Status(200)
	return c.JSON(context)
}
