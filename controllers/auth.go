package controllers

import (
	"fmt"
	"marketplace/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type LoginResponse struct {
	ID        uint             `json:"id"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	Name      string           `json:"name"`
	Email     string           `json:"email"`
	Role      string           `json:"role"`
	Addresses []models.Address `json:"addresses"`
	// Token     string           `json:"token"`
}

func LoginSeller(c *fiber.Ctx) error {
	var seller models.Seller
	err := c.BodyParser(&seller)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	sellerFromDB, err := models.FindSellerByEmail(seller.Email)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": fiber.StatusNotFound, "message": "Email is not found"})
	}

	loginResponse := LoginResponse{
		ID:        sellerFromDB.ID,
		CreatedAt: sellerFromDB.CreatedAt,
		UpdatedAt: sellerFromDB.UpdatedAt,
		Name:      sellerFromDB.Name,
		Role:      sellerFromDB.Role,
		Email:     sellerFromDB.Email,
		Addresses: sellerFromDB.Addresses,
		// Token:     token,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "ok",
		"data":   loginResponse,
	})
}

func RegisterSeller(c *fiber.Ctx) error {
	var seller models.Seller
	err := c.BodyParser(&seller)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	err = models.CreateSeller(&seller)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"code":    fiber.StatusConflict,
			"message": fmt.Sprintf("email %v has already been used", seller.Email),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"status":  "created",
		"message": "Your account has successfully registered",
	})
}
