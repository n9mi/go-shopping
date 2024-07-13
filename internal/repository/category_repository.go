package repository

import "go-shopping/internal/entity"

type CategoryRepository struct {
	Repository Repository[entity.Category]
}

func NewCategoryRepository() *CategoryRepository {
	return new(CategoryRepository)
}
