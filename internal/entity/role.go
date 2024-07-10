package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	DisplayName string
	Code        string  `gorm:"unique"`
	Users       []*User `gorm:"many2many:user_roles"`
}
