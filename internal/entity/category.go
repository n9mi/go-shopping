package entity

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	Name     string
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
