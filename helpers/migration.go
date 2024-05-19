package helpers

import (
	"marketplace/configs"
	"marketplace/models"
)

func Migration() {
	configs.DB.AutoMigrate(&models.Seller{}, &models.Address{}, &models.Product{}, &models.Category{})
}
