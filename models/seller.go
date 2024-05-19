package models

import (
	"database/sql"
	"marketplace/configs"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Seller struct {
	ID          uint           `json:"id" gorm:"primaryKey;column:id"`
	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Name        string         `json:"name" gorm:"column:name;not null"`
	Email       string         `json:"email" gorm:"column:email;not null"`
	PhoneNumber string         `json:"phone_number" gorm:"column:phone_number;not null"`
	StoreName   string         `json:"store_name" gorm:"column:store_name;not null"`
	Password    string         `json:"password" gorm:"column:password;not null"`
	Role        string         `json:"role" gorm:"column:role;not null"`
	Image       sql.NullString `json:"image" gorm:"column:image;default:null"`
	Gender      sql.NullString `json:"gender" gorm:"column:gender;default:null"`
	Birth       sql.NullTime   `json:"birth" gorm:"column:birth;default:null"`
	Addresses   []Address      `json:"addresses" gorm:"foreignKey:user_id;references:id"`
}

func (s *Seller) TableName() string {
	return "sellers"
}

func FindAllSellers() ([]*Seller, error) {
	var sellers []*Seller
	err := configs.DB.Find(&sellers).Error
	return sellers, err
}

func FindSellerByID(id int) (*Seller, error) {
	var seller Seller
	err := configs.DB.Take(&seller, "id = ?", id).Error
	return &seller, err
}

func FindSellerByEmail(email string) (*Seller, error) {
	var seller Seller
	result := configs.DB.Where("email = ?", email).Take(&seller)
	return &seller, result.Error
}

func CreateSeller(s *Seller) error {
	err := configs.DB.Create(&s).Error
	return err
}

func UpdateSeller(id int, seller *Seller) error {
	result := configs.DB.Model(&Seller{}).Where("id = ?", id).Updates(seller)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	return result.Error
}

func DeleteSeller(id int) error {
	result := configs.DB.Delete(&Seller{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return result.Error
}
