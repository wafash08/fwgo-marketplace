package models

import "time"

type Address struct {
	ID                   uint      `json:"id" gorm:"primaryKey;column:id"`
	CreatedAt            time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt            time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Type                 string    `json:"type" gorm:"column:type"`
	RecipientName        string    `json:"recipient_name" gorm:"column:recipient_name"`
	RecipientPhoneNumber string    `json:"recipient_phone_number" gorm:"column:recipient_phone_number"`
	Address              string    `json:"address" gorm:"column:address"`
	PostalCode           string    `json:"postal_code" gorm:"column:postal_code"`
	City                 string    `json:"city" gorm:"column:city"`
	Primary              bool      `json:"primary" gorm:"column:primary"`
	UserId               uint      `gorm:"column:user_id"`
	User                 Seller    `gorm:"foreignKey:user_id;references:id"`
}

func (a *Address) TableName() string {
	return "addresses"
}
