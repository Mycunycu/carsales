package config

import (
	"carsalesuser/logger"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	DEVELOPMENT = "development"
	PRODUCTION  = "production"
)

type Config struct {
	Environment string `yaml:"environment" env-default:"development"`
	RPCServer   struct {
		Port string `yaml:"port" env-default:"8081"`
	} `yaml:"rpcserver" env-required:"true"`
	PgConnection string `yaml:"pg_connection" env-required:"true"`
	PgMigration  string `yaml:"pg_migration" env-required:"true" env-default:"file://internal/store/postgres/migrations"`
}

var cfg *Config

func Get() *Config {
	return cfg
}

func init() {
	var once sync.Once

	once.Do(func() {
		logger := logger.Get()
		defer logger.Sync()

		logger.Info("Read application config")

		cfg = &Config{}

		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			help, _ := cleanenv.GetDescription(cfg, nil)
			logger.Info(help)
			logger.Fatal(err.Error())
		}
	})
}
