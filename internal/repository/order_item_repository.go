package repository

import "go-shopping/internal/entity"

type OrderItemRepository struct {
	Repository Repository[entity.OrderItem]
}

func NewOrderItemRepository() *OrderItemRepository {
	return new(OrderItemRepository)
}
