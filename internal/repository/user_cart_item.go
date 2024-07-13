package repository

import "go-shopping/internal/entity"

type UserCartItemRepository struct {
	Repository Repository[entity.UserCartItem]
}

func NewUserCartItemRepository() *UserCartItemRepository {
	return new(UserCartItemRepository)
}
