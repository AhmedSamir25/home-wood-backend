package controller

import (
	"homewood/database"
	"homewood/function"
	"homewood/model"
	"time"

	"log"
	"math/rand"

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

	if record.Email == "" || record.Password == "" {
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

func SendToken(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Token sent successfully",
	}

	credentials := new(model.User)
	if err := c.Bind().Body(credentials); err != nil {
		log.Printf("Error parsing request body: %v", err)
		context["statusText"] = "bad"
		context["msg"] = "invalid request"
		c.Status(fiber.StatusBadRequest)
		return c.JSON(context)
	}

	var user model.User
	result := database.DbConn.Where("email = ?", credentials.Email).First(&user)
	if result.Error != nil {
		context["statusText"] = "bad"
		context["msg"] = "email not found"
		c.Status(fiber.StatusNotFound)
		return c.JSON(context)
	}

	rand.Seed(time.Now().UnixNano())
	token := rand.Intn(900000) + 100000

	function.SendMail(user.Email, token)

	record := model.PasswordReset{
		Email:      user.Email,
		ResetToken: token,
	}
	if err := database.DbConn.Create(&record).Error; err != nil {
		log.Printf("Error saving reset token: %v", err)
		context["statusText"] = "bad"
		context["msg"] = "error in saving reset token"
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(context)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(context)
}

func CheckToken(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "check token successful",
	}
	credentials := new(model.PasswordReset)
	if err := c.Bind().Body(credentials); err != nil {
		log.Printf("Error parsing request body: %v", err)
		context["statusText"] = "bad"
		context["msg"] = "invalid request"
		c.Status(fiber.StatusBadRequest)
		return c.JSON(context)
	}
	var passwordReset model.PasswordReset
	result := database.DbConn.Where("email = ? AND reset_token = ?", credentials.Email, credentials.ResetToken).First(&passwordReset)
	if result.Error != nil {
		context["statusText"] = "bad"
		context["msg"] = "email or token is invalid"
		c.Status(400)
		return c.JSON(context)
	}
	c.Status(200)
	return c.JSON(context)
}
func ResetAndUpdatePassword(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Password reset and update successful",
	}

	credentials := new(model.PasswordReset)
	if err := c.Bind().Body(credentials); err != nil {
		log.Printf("Error parsing request body: %v", err)
		context["statusText"] = "bad"
		context["msg"] = "Invalid request"
		c.Status(fiber.StatusBadRequest)
		return c.JSON(context)
	}
	record := new(model.User)
	if err := c.Bind().Body(record); err != nil {
		log.Printf("Error parsing request body: %v", err)
		context["statusText"] = "bad"
		context["msg"] = "Invalid request"
		c.Status(fiber.StatusBadRequest)
		return c.JSON(context)
	}

	var passwordReset model.PasswordReset
	result := database.DbConn.Where("email = ? AND reset_token = ?", credentials.Email, credentials.ResetToken).First(&passwordReset)
	if result.Error != nil {
		context["statusText"] = "bad"
		context["msg"] = "Invalid email or token"
		c.Status(fiber.StatusBadRequest)
		return c.JSON(context)
	}

	result = database.DbConn.Delete(&passwordReset)
	if result.Error != nil {
		context["statusText"] = "bad"
		context["msg"] = "Error deleting reset token"
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(context)
	}

	var user model.User
	result = database.DbConn.Where("email = ?", record.Email).First(&user)
	if result.Error != nil {
		context["statusText"] = "bad"
		context["msg"] = "Invalid email"
		c.Status(fiber.StatusBadRequest)
		return c.JSON(context)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(record.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		context["statusText"] = "bad"
		context["msg"] = "Error hashing password"
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(context)
	}

	user.Password = string(hashedPassword)
	updateResult := database.DbConn.Save(&user)
	if updateResult.Error != nil {
		log.Println("Error saving new password")
		context["statusText"] = "bad"
		context["msg"] = "Error updating password"
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(context)
	}

	c.Status(fiber.StatusOK)
	return c.JSON(context)
}

func TestT(c fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Update Password successful",
	}
	return c.JSON(context)
}
