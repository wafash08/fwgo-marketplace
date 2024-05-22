package models

import (
	"marketplace/src/configs"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Address struct {
	ID                   uint      `json:"id" gorm:"primaryKey;column:id"`
	CreatedAt            time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Type                 string    `json:"type" gorm:"column:type" validate:"oneof=home office"`
	RecipientName        string    `json:"recipient_name" gorm:"column:recipient_name"`
	RecipientPhoneNumber string    `json:"recipient_phone_number" gorm:"column:recipient_phone_number"`
	Address              string    `json:"address" gorm:"column:address"`
	PostalCode           string    `json:"postal_code" gorm:"column:postal_code"`
	City                 string    `json:"city" gorm:"column:city"`
	Primary              bool      `json:"primary" gorm:"column:primary"`
	UserId               uint      `json:"user_id"`
}

func (a *Address) TableName() string {
	return "addresses"
}

func FindAllAddresses() ([]*Address, error) {
	var addresses []*Address
	err := configs.DB.Preload("User").Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func FindAddressByID(id int) (*Address, error) {
	var address Address
	err := configs.DB.Preload("User").Take(&address, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func CreateAddress(a *Address) error {
	err := configs.DB.Create(&a).Error
	return err
}

func UpdateAddress(id int, address *Address) error {
	result := configs.DB.Model(&Address{}).Where("id = ?", id).Updates(address)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	return result.Error
}

func DeleteAddress(id int) error {
	result := configs.DB.Delete(&Address{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return result.Error
}
