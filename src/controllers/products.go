package controllers

import (
	"fmt"
	"marketplace/src/helpers"
	"marketplace/src/middlewares"
	"marketplace/src/models"
	"math"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

type Pagination struct {
	CurrentPage int     `json:"current_page"`
	Limit       int     `json:"limit"`
	TotalData   int64   `json:"total_data"`
	TotalPage   float64 `json:"total_page"`
}

func FindAllProducts(c *fiber.Ctx) error {
	// pagination
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
	// sort
	sort := c.Query("sort")
	if sort == "" {
		sort = "ASC"
	}
	sortby := c.Query("orderBy")
	if sortby == "" {
		sortby = "name"
	}
	sort = sortby + " " + strings.ToLower(sort)
	// search
	keyword := c.Query("search")
	products := models.FindAllProducts(sort, keyword, limit, offset)
	totalData := models.CountData()
	totalPage := math.Ceil(float64(totalData) / float64(limit))
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
	user := middlewares.UserLocals(c)
	role := user["role"].(string)

	fmt.Println("user > ", user)
	fmt.Println("role > ", role)

	if role != "seller" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"code":    fiber.StatusForbidden,
			"message": "You don't have authorization for creating product",
		})
	}

	fmt.Println("user >> ", user)
	var product models.Product
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	// product = helpers.XSSMiddleware(product)

	// var newProduct models.Product
	// mapstructure.Decode(product, &newProduct)

	errors := helpers.ValidateStruct(product)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
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
	var product map[string]interface{}
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	product = helpers.XSSMiddleware(product)

	var updatedProduct models.Product
	mapstructure.Decode(product, &updatedProduct)

	errors := helpers.ValidateStruct(updatedProduct)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	err = models.UpdateProduct(id, &updatedProduct)
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
