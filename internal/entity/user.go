package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Roles    []*Role `gorm:"many2many:user_roles"`
}
