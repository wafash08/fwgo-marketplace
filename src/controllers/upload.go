package controllers

import (
	"fmt"
	"marketplace/src/helpers"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(c *fiber.Ctx) error {
	// Ambil file dari form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": fmt.Sprintf("Uploading file has failed because: %s", err.Error()),
		})
	}

	// Validasi ukuran file (maksimal 2MB)
	maxFileSize := int64(2 << 20) // 2MB
	if err := helpers.SizeUploadValidation(file.Size, maxFileSize); err != nil {
		return err
	}

	// Baca sebagian dari file untuk validasi tipe
	fileHeader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": fmt.Sprintf("Opening file has filed because: %s", err.Error()),
		})
	}
	defer fileHeader.Close()

	buffer := make([]byte, 512)
	_, err = fileHeader.Read(buffer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": fmt.Sprintf("Reading file has filed because: %s", err.Error()),
		})
	}

	// Validasi tipe file
	validFileTypes := []string{"image/png", "image/jpeg", "image/jpg", "application/pdf"}
	err = helpers.TypeUploadValidation(buffer, validFileTypes)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "The file you uploaded is not supported. Try uploading file which type is one of png, jpeg, jpg, and pdf",
		})
	}

	// Simpan file di direktori lokal
	filePath := helpers.UploadFile(file)
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": fmt.Sprintf("Failed to save file because: %s", err.Error()),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  "ok",
		"message": fmt.Sprintf("Your file has successfully uploaded to %s.", filePath),
	})
}