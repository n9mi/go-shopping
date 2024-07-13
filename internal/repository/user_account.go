package repository

import "go-shopping/internal/entity"

type UserPaymentAccountRepository struct {
	Repository Repository[entity.UserPaymentAccount]
}

func NewUserPaymentAccountRepository() *UserPaymentAccountRepository {
	return new(UserPaymentAccountRepository)
}
