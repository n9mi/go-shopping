package repository

import "go-shopping/internal/entity"

type OrderRepository struct {
	Repository Repository[entity.Order]
}

func NewOrderRepository() *OrderRepository {
	return new(OrderRepository)
}
