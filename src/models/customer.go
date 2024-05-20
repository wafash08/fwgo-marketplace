package models

import (
	"marketplace/src/configs"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Customer struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name" validate:"required,min=2"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Role      string    `json:"role" validate:"oneof=seller customer,required"`
}

func (c *Customer) TableName() string {
	return "customers"
}

func FindAllCustomers() ([]*Customer, error) {
	var customers []*Customer
	err := configs.DB.Find(&customers).Error
	return customers, err
}

func FindCustomerByID(id int) (*Customer, error) {
	var customer Customer
	err := configs.DB.Take(&customer, "id = ?", id).Error
	return &customer, err
}

func FindCustomerByEmail(email string) (*Customer, error) {
	var customer Customer
	result := configs.DB.Where("email = ?", email).Take(&customer)
	return &customer, result.Error
}

func CreateCustomer(out *Customer) error {
	err := configs.DB.Create(&out).Error
	return err
}

func UpdateCustomer(id int, c *Customer) error {
	result := configs.DB.Model(&Customer{}).Where("id = ?", id).Updates(c)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}

	return result.Error
}

func DeleteCustomer(id int) error {
	result := configs.DB.Delete(&Customer{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return result.Error
}
