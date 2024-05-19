package controllers

import (
	"fmt"
	"marketplace/src/helpers"
	"marketplace/src/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginResponse struct {
	ID           uint             `json:"id"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
	Name         string           `json:"name"`
	Email        string           `json:"email"`
	Role         string           `json:"role"`
	Addresses    []models.Address `json:"addresses"`
	Token        string           `json:"token"`
	RefreshToken string           `json:"refresh_token"`
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
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"code": fiber.StatusNotFound, "message": fmt.Sprintf("Account with email %s is not found", seller.Email)})
	}

	err = bcrypt.CompareHashAndPassword([]byte(sellerFromDB.Password), []byte(seller.Password))
	if err != nil || seller.Password == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"code":    fiber.StatusUnauthorized,
			"message": "Email or password is wrong",
		})
	}

	secretKey := os.Getenv("SECRET_KEY")
	payload := map[string]interface{}{
		"email": seller.Email,
	}

	token, err := helpers.GenerateToken(secretKey, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.StatusInternalServerError, "message": "Failed to generate token"})
	}

	refreshToken, err := helpers.GenerateRefreshToken(secretKey, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.StatusInternalServerError, "message": "Failed to generate refresh token"})
	}

	loginResponse := LoginResponse{
		ID:           sellerFromDB.ID,
		CreatedAt:    sellerFromDB.CreatedAt,
		UpdatedAt:    sellerFromDB.UpdatedAt,
		Name:         sellerFromDB.Name,
		Role:         sellerFromDB.Role,
		Email:        sellerFromDB.Email,
		Addresses:    sellerFromDB.Addresses,
		Token:        token,
		RefreshToken: refreshToken,
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

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(seller.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Ups, an error has occured in our server",
		})
	}
	seller.Password = string(hashPassword)

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

func RefreshToken(c *fiber.Ctx) error {
	var input struct {
		RefreshToken string `json:"refresh_token"`
	}
	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid request body",
		})
	}

	secretKey := os.Getenv("SECRET_KEY")
	token, err := helpers.GenerateToken(secretKey, map[string]interface{}{"refreshToken": input.RefreshToken})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Couldn't generate access token",
		})
	}

	refreshToken, err := helpers.GenerateRefreshToken(secretKey, map[string]interface{}{"refreshToken": input.RefreshToken})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate refresh token",
		})
	}

	item := map[string]string{
		"token":         token,
		"refresh_token": refreshToken,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"status":  "ok",
		"data":    item,
		"message": "Refresh token has successfully re-created",
	})
}
