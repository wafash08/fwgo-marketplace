package controllers

import (
	"fmt"
	"marketplace/models"
	"math"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Pagination struct {
	CurrentPage int     `json:"current_page"`
	Limit       int     `json:"limit"`
	TotalData   int64   `json:"total_data"`
	TotalPage   float64 `json:"total_page"`
}

func FindAllProducts(c *fiber.Ctx) error {
	pageOld := c.Query("page")
	limitOld := c.Query("limit")
	page, _ := strconv.Atoi(pageOld)
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(limitOld)
	if limit == 0 {
		limit = 5
	}
	offset := (page - 1) * limit
	sort := c.Query("sort")
	if sort == "" {
		sort = "ASC"
	}
	sortby := c.Query("orderBy")
	if sortby == "" {
		sortby = "name"
	}
	sort = sortby + " " + strings.ToLower(sort)
	keyword := c.Query("search")
	products, err := models.FindAllProducts(sort, keyword, limit, offset)
	totalData := models.CountData()
	totalPage := math.Ceil(float64(totalData) / float64(limit))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Ups, unknown error has occured in our server.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   products,
		"pagination": Pagination{
			CurrentPage: page,
			Limit:       limit,
			TotalData:   totalData,
			TotalPage:   totalPage,
		},
	})
}

func FindProductById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := models.FindProductByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Product with id %d is not found", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   product,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	err = models.CreateProduct(&product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"status":  "created",
		"message": "Product has successfully created",
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	err = models.UpdateProduct(id, &product)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to update product with ID %d because there is no product with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  "ok",
		"message": fmt.Sprintf("Product with id %d has successfully updated.", id),
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to delete product with ID %d because there is no product with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  "ok",
		"message": fmt.Sprintf("Product with id has %d successfully deleted.", id),
	})
}
