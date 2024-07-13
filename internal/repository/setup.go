package repository

type RepositorySetup struct {
	RoleRepository               *RoleRepository
	UserRepository               *UserRepository
	CategoryRepository           *CategoryRepository
	ProductRepository            *ProductRepository
	PaymentMethodRepository      *PaymentMethodRepository
	UserPaymentAccountRepository *UserPaymentAccountRepository
	OrderRepository              *OrderRepository
	OrderItemRepository          *OrderItemRepository
	UserCartItemRepository       *UserCartItemRepository
}

func Setup() *RepositorySetup {
	return &RepositorySetup{
		RoleRepository:               NewRoleRepository(),
		UserRepository:               NewUserRepository(),
		CategoryRepository:           NewCategoryRepository(),
		ProductRepository:            NewProductRepository(),
		PaymentMethodRepository:      NewPaymentMethodRepository(),
		UserPaymentAccountRepository: NewUserPaymentAccountRepository(),
		OrderRepository:              NewOrderRepository(),
		OrderItemRepository:          NewOrderItemRepository(),
		UserCartItemRepository:       NewUserCartItemRepository(),
	}
}
