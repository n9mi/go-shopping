package repository

import (
	"go-shopping/internal/entity"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

func (r *UserRepository) ExistsByEmail(db *gorm.DB, email string) bool {
	var exists bool

	db.Model(&entity.User{}).
		Select("count(email) > 0").
		Where("email = ?", strings.ToLower(email)).
		Find(&exists)

	return exists
}
