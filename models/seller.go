package models

import (
	"database/sql"
	"time"
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
