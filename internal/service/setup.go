package service

import (
	"go-shopping/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ServiceSetup struct {
	AuthService AuthService
}

func Setup(viperCfg *viper.Viper, logger *logrus.Logger, db *gorm.DB,
	validator *validator.Validate, repositorySetup *repository.RepositorySetup) *ServiceSetup {
	return &ServiceSetup{
		AuthService: NewAuthService(viperCfg, logger, db, validator, repositorySetup.UserRepository),
	}
}
