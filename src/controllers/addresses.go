package controllers

import (
	"fmt"
	"marketplace/src/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func FindAllAddresses(c *fiber.Ctx) error {
	addresses, err := models.FindAllAddresses()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Ups, an unknown error has occured in our server",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   addresses,
	})
}

func FindAddressByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	address, err := models.FindAddressByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Address with id %d is not found", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   address,
	})
}

func CreateAddress(c *fiber.Ctx) error {
	var address models.Address
	err := c.BodyParser(&address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	err = models.CreateAddress(&address)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"status":  "created",
		"message": "Address has successfully created",
	})
}

func UpdateAddress(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var address models.Address
	err := c.BodyParser(&address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	err = models.UpdateAddress(id, &address)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to update address with ID %d because there is no address with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
	})
}

func DeleteAddress(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	err := models.DeleteAddress(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"code":    fiber.StatusNotFound,
			"message": fmt.Sprintf("Failed to delete address with ID %d because there is no address with such id", id),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
	})
}
