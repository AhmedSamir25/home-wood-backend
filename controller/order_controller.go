package controller

import (
	"homewood/common"
	"homewood/database"
	"homewood/model"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func AddOrder(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "order successfully",
		"statusText": "Ok",
	}
	record := new(model.Orders)
	db := database.DbConn
	if err := c.Bind().Body(record); err != nil {
		context["msg"] = "invalid request"
		context["statusText"] = "error"
		c.Status(400)
		return c.JSON(context)
	}
	result := db.Create(record)
	if result.Error != nil {
		context["msg"] = "An error occurred while placing an order"
		context["statusText"] = "error"
		return c.Status(400).JSON(context)
	}

	return c.Status(200).JSON(context)
}

func OrderUser(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "get all orders success",
		"statusText": "Ok",
	}
	db := database.DbConn

	pageStr := c.Params("pageId", "1")
	userId := c.Params("userId")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	perPage := c.Query("per_page", "10")
	sortOrder := c.Query("sort_order", "desc")
	record := []model.Products{}

	limit, err := strconv.ParseInt(perPage, 10, 64)
	if limit < 1 || limit > 100 {
		limit = 5
	}
	if err != nil {
		return c.Status(500).JSON("Invalid per_page option")
	}

	offset := (page - 1) * int(limit)
	db.Table("products").
		Select("*").
		Joins("JOIN orders AS odr ON odr.product_id = products.product_id").
		Where("odr.user_id = ?", userId).
		Order("products.product_id " + sortOrder).
		Offset(offset).
		Limit(int(limit)).
		Find(&record)

	pageInfo := calculatePagination(page == 1, len(record) == int(limit), int(limit), record, len(record) == int(limit))

	response := common.ResponseDTO{
		Success:    true,
		Data:       record,
		Pagination: pageInfo,
	}
	context["products"] = response
	return c.Status(200).JSON(context)
}
