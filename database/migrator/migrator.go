package migrator

import (
	"fmt"
	"go-shopping/internal/entity"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, logger *logrus.Logger) {
	if err := db.AutoMigrate(&entity.Role{}); err != nil {
		logger.Fatalf("failed to create roles table: %+v", err)
	}
	logger.Info("successfully create roles table")

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		logger.Fatalf("failed to create users table: %+v", err)
	}
	logger.Info("successfully create users table")

	if err := db.AutoMigrate(&entity.Category{}); err != nil {
		logger.Fatalf("failed to create categorys table: %+v", err)
	}
	logger.Info("successfully create categorys table")

	if err := db.AutoMigrate(&entity.Product{}); err != nil {
		logger.Fatalf("failed to create products table: %+v", err)
	}
	logger.Info("successfully create products table")

	if err := db.AutoMigrate(&entity.PaymentMethod{}); err != nil {
		logger.Fatalf("failed to create payment_methods table: %+v", err)
	}
	logger.Info("successfully create payment_methods table")

	if err := db.AutoMigrate(&entity.UserPaymentAccount{}); err != nil {
		logger.Fatalf("failed to create user_accounts table: %+v", err)
	}
	logger.Info("successfully create user_accounts table")

	if err := db.AutoMigrate(&entity.Order{}); err != nil {
		logger.Fatalf("failed to create orders table: %+v", err)
	}
	logger.Info("successfully create orders table")

	if err := db.AutoMigrate(&entity.OrderItem{}); err != nil {
		logger.Fatalf("failed to create order_items table: %+v", err)
	}
	logger.Info("successfully create order_items table")

	if err := db.AutoMigrate(&entity.UserCartItem{}); err != nil {
		logger.Fatalf("failed to create user_cart_items table: %+v", err)
	}
	logger.Info("successfully create user_cart_items table")
}

func Drop(db *gorm.DB, logger *logrus.Logger, viperCfg *viper.Viper) {
	if err := db.Migrator().DropTable(&entity.UserCartItem{}); err != nil {
		logger.Fatalf("failed to drop user_cart_items table: %+v", err)
	}
	logger.Info("successfully dropping user_cart_items table")

	if err := db.Migrator().DropTable(&entity.OrderItem{}); err != nil {
		logger.Fatalf("failed to drop order_items table: %+v", err)
	}
	logger.Info("successfully dropping order_items table")

	if err := db.Migrator().DropTable(&entity.Order{}); err != nil {
		logger.Fatalf("failed to drop orders table: %+v", err)
	}
	logger.Info("successfully dropping orders table")

	if err := db.Migrator().DropTable(&entity.UserPaymentAccount{}); err != nil {
		logger.Fatalf("failed to drop user_accounts table: %+v", err)
	}
	logger.Info("successfully dropping user_accounts table")

	if err := db.Migrator().DropTable(&entity.PaymentMethod{}); err != nil {
		logger.Fatalf("failed to drop payment_methods table: %+v", err)
	}
	logger.Info("successfully dropping payment_methods table")

	if err := db.Migrator().DropTable(&entity.Product{}); err != nil {
		logger.Fatalf("failed to drop products table: %+v", err)
	}
	logger.Info("successfully dropping products table")

	if err := db.Migrator().DropTable(&entity.Category{}); err != nil {
		logger.Fatalf("failed to drop categorys table: %+v", err)
	}
	logger.Info("successfully dropping categories table")

	if err := db.Migrator().DropTable(fmt.Sprintf("%s_%s_user_roles",
		viperCfg.GetString("PROJECT_ID"),
		viperCfg.GetString("ENV"))); err != nil {
		logger.Fatalf("failed to drop user_roles table: %+v", err)
	}
	logger.Info("successfully dropping user_roles table")

	if err := db.Migrator().DropTable(&entity.User{}); err != nil {
		logger.Fatalf("failed to drop users table: %+v", err)
	}
	logger.Info("successfully dropping users table")

	if err := db.Migrator().DropTable(&entity.Role{}); err != nil {
		logger.Fatalf("failed to drop roles table: %+v", err)
	}
	logger.Info("successfully dropping roles table")
}
