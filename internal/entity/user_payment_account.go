package entity

import "gorm.io/gorm"

type UserPaymentAccount struct {
	gorm.Model
	ID              string `gorm:"primaryKey"`
	PaymentMethodID uint64
	PaymentMethod   PaymentMethod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID          string
	User            User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	AccountNumber   string  `gorm:"unique"`
	Orders          []Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
