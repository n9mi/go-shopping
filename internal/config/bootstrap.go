package config

import (
	"go-shopping/database/migrator"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ConfigBootstrap struct {
	ViperCfg  *viper.Viper
	App       *fiber.App
	Logger    *logrus.Logger
	DB        *gorm.DB
	Validator *validator.Validate
}

func Bootstrap(cfgBootstrap *ConfigBootstrap) {
	migrator.Drop(cfgBootstrap.DB, cfgBootstrap.Logger)
	migrator.Create(cfgBootstrap.DB, cfgBootstrap.Logger)
}
