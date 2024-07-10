package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
	Roles    []*Role `gorm:"many2many:240708gos_user_roles"`
}

func (e *User) TableName() string {
	return "240708gos_users"
}
