package repository

import "go-shopping/internal/entity"

type RoleRepository struct {
	Repository Repository[entity.Role]
}
