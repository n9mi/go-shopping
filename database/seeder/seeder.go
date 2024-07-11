package seeder

import (
	"fmt"
	"go-shopping/internal/entity"
	"go-shopping/internal/utils"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB, logger *logrus.Logger) {
	roles := seedRoles(db, logger)
	users := seedUsers(db, logger, roles)
	categories := seedCategories(db, logger)
	products := seedProducts(db, logger, categories)
	seedUserCartItems(db, logger, users, products)
	paymentMethods := seedPaymentMethod(db, logger)
	seedUserPaymentAccounts(db, logger, users, paymentMethods)
}

func seedRoles(db *gorm.DB, logger *logrus.Logger) []entity.Role {
	roles := []entity.Role{
		{
			DisplayName: "Customer",
			Code:        "customer",
		},
	}

	if err := db.Create(roles).Error; err != nil {
		logger.Fatalf("failed to seed roles table : %+v", err)
	}

	return roles
}

func seedUsers(db *gorm.DB, logger *logrus.Logger, roles []entity.Role) []entity.User {
	var users []entity.User
	numOfUsers := 5
	userPwd, _ := utils.GeneratePassword("password")

	for _, r := range roles {
		for i := 1; i <= numOfUsers; i++ {
			users = append(users, entity.User{
				ID:       uuid.NewString(),
				Name:     fmt.Sprintf("%s %d", r.DisplayName, i),
				Email:    fmt.Sprintf("%s%d@test.com", r.Code, i),
				Password: userPwd,
				Roles:    []*entity.Role{&r},
			})
		}
	}

	if err := db.Create(users).Error; err != nil {
		logger.Fatalf("failed to seed users table : %+v", err)
	}

	return users
}

func seedCategories(db *gorm.DB, logger *logrus.Logger) []entity.Category {
	var categories []entity.Category
	numOfCategories := 5

	for i := 1; i <= numOfCategories; i++ {
		categories = append(categories, entity.Category{
			Name: fmt.Sprintf("Category %d", i),
		})
	}

	if err := db.Create(categories).Error; err != nil {
		logger.Fatalf("failed to seed categories table : %+v", err)
	}

	return categories
}

func seedProducts(db *gorm.DB, logger *logrus.Logger, categories []entity.Category) []entity.Product {
	var products []entity.Product
	numOfProductPerCategory := 5
	counter := 0

	for _, c := range categories {
		for i := 0; i < numOfProductPerCategory; i++ {
			counter += 1

			products = append(products, entity.Product{
				Name:       fmt.Sprintf("Product %d", counter),
				PriceIDR:   float64(utils.RandomNumberWithRange(1, 100)) * 1000,
				CategoryID: c.ID,
			})
		}
	}

	if err := db.Create(products).Error; err != nil {
		logger.Fatalf("failed to seed products table : %+v", err)
	}

	return products
}

func seedUserCartItems(db *gorm.DB, logger *logrus.Logger, users []entity.User, products []entity.Product) {
	var userCartItems []entity.UserCartItem
	numOfProducts := len(products)

	for _, u := range users {
		numProdPerCart := utils.RandomNumberWithRange(1, 3)
		counter := 0
		var productIncluded []int

		for counter <= numProdPerCart {
			randProdIx := utils.RandomNumberWithRange(1, numOfProducts)

			if !slices.Contains(productIncluded, randProdIx) {
				productIncluded = append(productIncluded, randProdIx)
				userCartItems = append(userCartItems, entity.UserCartItem{
					UserID:    u.ID,
					ProductID: products[randProdIx].ID,
					Quantity:  utils.RandomNumberWithRange(1, 10),
				})
				counter += 1
			}
		}
	}

	if err := db.Create(userCartItems).Error; err != nil {
		logger.Fatalf("failed to seed user_cart_items table : %+v", err)
	}
}

func seedPaymentMethod(db *gorm.DB, logger *logrus.Logger) []entity.PaymentMethod {
	var paymentMethods []entity.PaymentMethod
	numOfPaymentMethods := 5

	for i := 1; i <= numOfPaymentMethods; i++ {
		paymentMethods = append(paymentMethods, entity.PaymentMethod{
			Name: fmt.Sprintf("Payment method %d", i),
		})
	}

	if err := db.Create(paymentMethods).Error; err != nil {
		logger.Fatalf("failed to seed payment_methods table : %+v", err)
	}

	return paymentMethods
}

func seedUserPaymentAccounts(db *gorm.DB, logger *logrus.Logger, users []entity.User, paymentMethods []entity.PaymentMethod) []entity.UserPaymentAccount {
	var userPaymentAccounts []entity.UserPaymentAccount
	numOfPaymentMethods := len(paymentMethods)

	for i, u := range users {
		userPaymentAccounts = append(userPaymentAccounts, entity.UserPaymentAccount{
			ID:              uuid.NewString(),
			PaymentMethodID: paymentMethods[uint64(utils.RandomNumberWithRange(1, numOfPaymentMethods))].ID,
			UserID:          u.ID,
			AccountNumber:   strings.Repeat(strconv.Itoa(i), 4),
		})
	}

	if err := db.Create(userPaymentAccounts).Error; err != nil {
		logger.Fatalf("failed to seed user_payment_accounts table : %+v", err)
	}

	return userPaymentAccounts
}
