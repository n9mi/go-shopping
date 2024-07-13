package config

import (
	"go-shopping/database/migrator"
	"go-shopping/database/seeder"
	"go-shopping/internal/delivery/http/controller"
	"go-shopping/internal/delivery/http/route"
	"go-shopping/internal/repository"
	"go-shopping/internal/service"

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
	repositorySetup := repository.Setup()
	serviceSetup := service.Setup(
		cfgBootstrap.ViperCfg,
		cfgBootstrap.Logger,
		cfgBootstrap.DB,
		cfgBootstrap.Validator,
		repositorySetup)
	controllerSetup := controller.Setup(
		cfgBootstrap.Logger,
		serviceSetup)

	routeConfig := route.RouteConfig{
		App:             cfgBootstrap.App,
		ControllerSetup: controllerSetup,
	}
	routeConfig.Setup()

	migrator.Drop(cfgBootstrap.DB, cfgBootstrap.Logger, cfgBootstrap.ViperCfg)
	migrator.Create(cfgBootstrap.DB, cfgBootstrap.Logger)
	seeder.Seed(cfgBootstrap.DB, cfgBootstrap.Logger)
}
