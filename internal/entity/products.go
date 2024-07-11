package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	Name       string
	PriceIDR   float64 `gorm:"column:price_idr"`
	CategoryID uint64
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
