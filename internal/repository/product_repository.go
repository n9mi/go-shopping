package repository

import "go-shopping/internal/entity"

type ProductRepository struct {
	Repository Repository[entity.Product]
}

func NewProductRepository() *ProductRepository {
	return new(ProductRepository)
}
