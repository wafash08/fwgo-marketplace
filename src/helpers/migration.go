package helpers

import (
	"marketplace/src/configs"
	"marketplace/src/models"
)

func Migration() {
	configs.DB.AutoMigrate(
		&models.Seller{},
		&models.Address{},
		&models.Product{},
		&models.Category{},
		&models.Customer{},
	)
}
