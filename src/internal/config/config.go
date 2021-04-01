package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Port         string
	Env          string
	PgConnStr    string
	MongoConnStr string
	DbName       string
}

var config Config

func Init() (*Config, error) {
	var once sync.Once

	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			logrus.Fatal("Error loading .env file")
		}

		env := getStringEnv("ENV", "")
		if env == "" {
			env = "dev"
		}

		config.Env = env
		config.Port = getStringEnv("PORT", "")
		config.PgConnStr = getStringEnv("PG_CONNECTION", "")
		config.MongoConnStr = getStringEnv("MONGO_CONNECTION", "")
		config.DbName = getStringEnv("DB_NAME", "")
	})

	return &config, nil
}

func getStringEnv(key string, defValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defValue
	}

	return value
}
