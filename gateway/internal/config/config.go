package config

import (
	"fmt"
	"sync"

	"github.com/Mycunycu/carsales/gateway/pkg/logger"

	"github.com/ilyakaznacheev/cleanenv"
)

type Environment string

const (
	Dev  Environment = "development"
	Stag Environment = "staging"
	Prod Environment = "production"
)

type Config struct {
	Environment string `yaml:"environment" env-default:"development"`
	HTTPServer  struct {
		Type   string `yaml:"type" env-default:"port"`
		Domain string `yaml:"domain" env-default:"localhost"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"httpserver"`
	JWT struct {
		Secret string `yaml:"secret" env-required:"true"`
	}
	UserService struct {
		Port string `yaml:"port" env-required:"true"`
	} `yaml:"userservice" env-required:"true"`
}

var cfg *Config

func GetConfig() *Config {
	return cfg
}

func Init() error {
	var once sync.Once
	var err error

	once.Do(func() {
		logger := logger.GetLogger()
		defer logger.Sync()

		cfg = &Config{}

		logger.Info("read application config")
		if err = cleanenv.ReadConfig("config.yml", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			logger.Info(help)
			err = fmt.Errorf("cleanenv.ReadConfig %w", err)
		}
	})

	return err
}
