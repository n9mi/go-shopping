package migrator

import (
	"go-shopping/internal/entity"

	"github.com/sirupsen/logrus"
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
}

func Drop(db *gorm.DB, logger *logrus.Logger) {
	if err := db.Migrator().DropTable(&entity.User{}); err != nil {
		logger.Fatalf("failed to drop users table: %+v", err)
	}
	logger.Info("successfully dropping users table")

	if err := db.Migrator().DropTable(&entity.Role{}); err != nil {
		logger.Fatalf("failed to drop roles table: %+v", err)
	}
	logger.Info("successfully dropping roles table")
}
