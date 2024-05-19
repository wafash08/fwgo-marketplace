package models

import (
	"marketplace/src/configs"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Category struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" gorm:"not null"`
	Color     string    `json:"color" gorm:"not null"`
	Image     string    `json:"image" gorm:"not null"`
	Products  []Product `json:"products" gorm:"foreignKey:category_id;references:id"`
}

func (c *Category) TableName() string {
	return "categories"
}

type ProductResponse struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Brand       string    `json:"brand"`
	Rating      int       `json:"rating"`
	Price       float64   `json:"price"`
	Color       string    `json:"color"`
	Size        int       `json:"size"`
	Quantity    int       `json:"quantity"`
	Image       string    `json:"image"`
	Condition   string    `json:"condition"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"category_id"`
}

func FindAllCategories(sort, name string) ([]*Category, error) {
	var categories []*Category
	name = "%" + name + "%"
	err := configs.DB.Preload("Products").Order(sort).Where("name ILIKE ?", name).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func FindCategoryByID(id int) (*Category, error) {
	var category Category
	err := configs.DB.Preload("Products").Take(&category, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func CreateCategory(c *Category) error {
	err := configs.DB.Create(&c).Error
	return err
}

func UpdateCategory(id int, category *Category) error {
	result := configs.DB.Model(&Category{}).Where("id = ?", id).Updates(category)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	return result.Error
}

func DeleteCategory(id int) error {
	result := configs.DB.Delete(&Category{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return result.Error
}
