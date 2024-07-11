package repository

import "go-shopping/internal/entity"

type UserAccountRepository struct {
	Repository Repository[entity.UserPaymentAccount]
}
