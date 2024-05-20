package helpers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateStruct(param any) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			element := ErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			}
			errors = append(errors, &element)
		}
	}
	return errors
}

func SizeUploadValidation(fileSize int64, maxFileSize int64) error {
	if fileSize > maxFileSize {
		return fiber.NewError(fiber.StatusRequestEntityTooLarge, "Ukuran file melebihi 2MB")
	}
	return nil
}

func TypeUploadValidation(buffer []byte, validFileTypes []string) error {
	fileType := http.DetectContentType(buffer)
	if !isValidFileType(validFileTypes, fileType) {
		return fiber.NewError(fiber.StatusBadRequest, "Tipe file tidak valid. Hanya png, jpg, jpeg, dan pdf yang diperbolehkan.")
	}
	return nil
}

func isValidFileType(validFileTypes []string, fileType string) bool {
	for _, validType := range validFileTypes {
		if validType == fileType {
			return true
		}
	}
	return false
}

var (
	NotIncludeSpace           = regexp.MustCompile(`\s`)
	MustIncludeOneUppercase   = regexp.MustCompile(`[A-Z]`)
	MustIncludeOneNumber      = regexp.MustCompile(`[0-9]`)
	MustIncludeOneSpecialChar = regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`)
)

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password should contain minimal 8 characters")
	} else if NotIncludeSpace.MatchString(password) {
		return fmt.Errorf("password should not contain space")
	} else if !MustIncludeOneUppercase.MatchString(password) {
		return fmt.Errorf("password should contain one uppercase")
	} else if !MustIncludeOneNumber.MatchString(password) {
		return fmt.Errorf("password should contain one number")
	} else if !MustIncludeOneSpecialChar.MatchString(password) {
		return fmt.Errorf("password should contain one special character")
	} else {
		return nil
	}
}
