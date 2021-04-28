package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

const (
	DEV_ENV  = "dev"
	PROD_ENV = "prod"
)

type Config struct {
	Port         string
	Env          string
	PgConnStr    string
	MongoConnStr string
	DbName       string
}

var cfg *Config

func Get() *Config {
	return cfg
}

func Init() {
	var once sync.Once

	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}

		cfg = &Config{}

		cfg.Env = getVar("ENV", DEV_ENV)
		cfg.Port = getVar("PORT", "")
		cfg.PgConnStr = getVar("PG_CONNECTION", "")
		cfg.MongoConnStr = getVar("MONGO_CONNECTION", "")
		cfg.DbName = getVar("DB_NAME", "")
	})
}

func getVar(key string, defValue string) string {
	value, isExist := os.LookupEnv(key)

	if isExist {
		return value
	}

	if defValue == "" {
		log.Fatalf("Environment variable %s not defined", key)
	}

	return defValue
}
