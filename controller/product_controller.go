package controller

import (
	"homewood/common"
	"homewood/database"
	"homewood/helpers"
	"homewood/model"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

func GetAllProducts(c fiber.Ctx) error {
	context := fiber.Map{
		"msg":        "get all products success",
		"statusText": "Ok",
	}
	db := database.DbConn

	pageStr := c.Params("pageid", "1")
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

	db.Order("product_id " + sortOrder).Offset(offset).Limit(int(limit)).Find(&record)

	if len(record) == 0 {
		context["msg"] = "No products found"
		context["statusText"] = "error"
		return c.Status(400).JSON(context)
	}

	pageInfo := calculatePagination(page == 1, len(record) == int(limit), int(limit), record, len(record) == int(limit))

	response := common.ResponseDTO{
		Success:    true,
		Data:       record,
		Pagination: pageInfo,
	}
	context["products"] = response
	return c.Status(200).JSON(context)
}

func calculatePagination(isFirstPage bool, hasPagination bool, _ int, record []model.Products, pointsNext bool) helpers.PaginationInfo {
	pagination := helpers.PaginationInfo{}
	var nextCur, prevCur helpers.Cursor

	if isFirstPage {
		if hasPagination {
			nextCur = helpers.CreateCursor(record[len(record)-1].ProductId, true)
			pagination = helpers.GeneratePager(nextCur, nil)
		}
	} else {
		if pointsNext {
			if hasPagination {
				nextCur = helpers.CreateCursor(record[len(record)-1].ProductId, true)
			}
			prevCur = helpers.CreateCursor(record[0].ProductId, false)
			pagination = helpers.GeneratePager(nextCur, prevCur)
		} else {
			if len(record) > 0 {
				nextCur = helpers.CreateCursor(record[len(record)-1].ProductId, true)
				prevCur = helpers.CreateCursor(record[0].ProductId, false)
				pagination = helpers.GeneratePager(nextCur, prevCur)
			}
		}
	}
	return pagination
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
		return c.Status(400).JSON(context)
	}

	db := database.DbConn

	result := db.Create(record)
	if result.Error != nil {
		context["msg"] = "error when adding product"
		context["statusText"] = "error"
		return c.Status(400).JSON(context)
	}

	return c.Status(200).JSON(context)
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

func GetProductPyCategories() {
	log.Println("GetProductPyCategories")
}
