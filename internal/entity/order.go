package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID                   string `gorm:"primaryKey"`
	TotalPrice           int
	Status               int8
	User                 User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID               string
	UserPaymentAccount   UserPaymentAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserPaymentAccountID string
	OrderItems           []OrderItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
