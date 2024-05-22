package controllers

import (
	"fmt"
	"marketplace/src/helpers"
	"marketplace/src/models"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mitchellh/mapstructure"
)

func FindAllCategories(c *fiber.Ctx) error {
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
	categories, err := models.FindAllCategories(sort, keyword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Ups, an unknown error has occured in our server",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   categories,
	})
}

func FindCategoryByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	category, err := models.FindCategoryByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": "Category is not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   category,
	})
}

func CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	errors := helpers.ValidateStruct(category)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	err = models.CreateCategory(&category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"status":  "created",
		"message": "Category has successfully created",
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var category map[string]interface{}
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	category = helpers.XSSMiddleware(category)

	// Convert map to Category model using mapstructure
	var updatedCategory models.Category
	mapstructure.Decode(category, &updatedCategory)

	errors := helpers.ValidateStruct(updatedCategory)
	if len(errors) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	err = models.UpdateCategory(id, &updatedCategory)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to update category with ID %d because there is no category with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteCategory(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to delete category with ID %d because there is no category with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
	})
}
