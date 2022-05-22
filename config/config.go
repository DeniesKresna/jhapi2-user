package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigV1 struct {
	Application *ApplicationConfig `yaml:"application"`
	Database    *DatabaseConfig    `yaml:"database"`
}

type ApplicationConfig struct {
	Environment *string `yaml:"environment"`
	Name        *string `yaml:"name"`
	Url         *string `yaml:"url"`
	Debug       *bool   `yaml:"debug"`
}

type DatabaseConfig struct {
	DSN *string `yaml:"dsn"`
}

var (
	cfg        ConfigV1
	gormClient *gorm.DB
)

func Get() ConfigV1 {
	return cfg
}

func GetDB() *gorm.DB {
	return gormClient
}

func ProvideConfig() {
	env := os.Getenv("ENV")
	err := cleanenv.ReadConfig("config/files/"+env+".config.yml", &cfg)
	if err != nil {
		log.Panic().Err(err).Msg("Fail create env config from file")
	}
	log.Info().Msg("Config initializing from file is done")
}

func ProvideDB() {
	dsn := os.Getenv("DSN")
	var err error

	gormClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("fail to initialize database")
	}

	sqlDB, err := gormClient.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Info().Msg("database connected")
	if *Get().Application.Debug {
		log.Debug().Msg("enable debug db")
		gormClient = gormClient.Debug()
	}
}
