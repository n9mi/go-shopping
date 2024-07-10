package main

import (
	"fmt"
	"go-shopping/internal/config"
)

func main() {
	viperCfg, err := config.NewViperConfig()
	if err != nil {
		panic(err)
	}

	app := config.NewFiber(viperCfg)
	log := config.NewLogrus(viperCfg)
	db := config.NewDatabase(viperCfg, log)
	validator := config.NewValidator()

	config.Bootstrap(&config.ConfigBootstrap{
		App:       app,
		Logger:    log,
		DB:        db,
		Validator: validator,
	})

	port := viperCfg.GetInt("APP_PORT")
	if port == 0 {
		port = viperCfg.GetInt("PORT")
	}

	app.Listen(fmt.Sprintf(":%d", port))
}
