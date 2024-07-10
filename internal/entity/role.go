package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	DisplayName string
	Code        string  `gorm:"unique"`
	Users       []*User `gorm:"many2many:240708gos_user_roles"`
}

func (e *Role) TableName() string {
	return "240708gos_roles"
}
