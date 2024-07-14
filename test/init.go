package test

import (
	"go-shopping/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	viperCfg *viper.Viper
	log      *logrus.Logger
	app      *fiber.App
	db       *gorm.DB
	validate *validator.Validate
	err      error
)

func init() {
	viperCfg, err = config.NewViperConfig()
	if err != nil {
		panic(err)
	}

	app = config.NewFiber(viperCfg)
	log = config.NewLogrus(viperCfg)
	db = config.NewDatabase(viperCfg, log)
	validate = config.NewValidator()

	config.Bootstrap(&config.ConfigBootstrap{
		ViperCfg:  viperCfg,
		App:       app,
		Logger:    log,
		DB:        db,
		Validator: validate,
	})
}
