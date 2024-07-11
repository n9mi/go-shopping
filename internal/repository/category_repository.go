package repository

import "go-shopping/internal/entity"

type CategoryRepository struct {
	Repository Repository[entity.Category]
}
