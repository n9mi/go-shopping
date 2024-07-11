package repository

import "go-shopping/internal/entity"

type UserRepository struct {
	Repository Repository[entity.User]
}
