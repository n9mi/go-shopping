package entity

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	ID                  uint64 `gorm:"primaryKey;autoIncrement"`
	Name                string
	UserPaymentAccounts []UserPaymentAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
