package entity

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Order     Order  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderID   string
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID uint64
	Quantity  int
}
