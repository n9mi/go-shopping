package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type logrusWriter struct {
	Logger *logrus.Logger
}

func (w *logrusWriter) Printf(message string, args ...interface{}) {
	w.Logger.Tracef(message, args...)
}

func NewDatabase(viperCfg *viper.Viper, log *logrus.Logger) *gorm.DB {
	host := viperCfg.GetString("DB_HOST")
	name := viperCfg.GetString("DB_NAME")
	port := viperCfg.GetInt("DB_PORT")
	user := viperCfg.GetString("DB_USER")
	password := viperCfg.GetString("DB_PASSWORD")
	ssl := viperCfg.GetString("DB_SSL_MODE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		host,
		user,
		password,
		name,
		port,
		ssl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             5 * time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	})
	if err != nil {
		log.Fatalf("failed to connect into database %+v \n", err)
	}

	return db
}
