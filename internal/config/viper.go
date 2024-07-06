package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

func NewViperConfig() (*viper.Viper, error) {
	viperCfg := viper.New()

	if _, err := os.Stat(".env"); errors.Is(err, os.ErrNotExist) {
		viperCfg.AutomaticEnv()
	} else {
		viperCfg.SetConfigType("env")
		viperCfg.SetConfigFile(".env")
		viperCfg.AddConfigPath("./")

		if err != nil {
			return nil, err
		}
	}

	return viperCfg, nil
}
