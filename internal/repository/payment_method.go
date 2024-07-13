package repository

import "go-shopping/internal/entity"

type PaymentMethodRepository struct {
	Repository Repository[entity.PaymentMethod]
}

func NewPaymentMethodRepository() *PaymentMethodRepository {
	return new(PaymentMethodRepository)
}
