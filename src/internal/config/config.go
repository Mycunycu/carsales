package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Env          string
	PgConnStr    string
	MongoConnStr string
	DbName       string
}

var config Config

func GetConfig() *Config {
	config.Env = getStringEnv("ENV", "dev")
	config.Port = getStringEnv("PORT", "")
	config.PgConnStr = getStringEnv("PG_CONNECTION", "")
	config.MongoConnStr = getStringEnv("MONGO_CONNECTION", "")
	config.DbName = getStringEnv("DB_NAME", "")

	return &config
}

func Init() *Config {
	var once sync.Once

	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	})

	config := GetConfig()

	return config
}

func getStringEnv(key string, defValue string) string {
	value, isExist := os.LookupEnv(key)

	if isExist {
		return value
	}

	if defValue == "" {
		log.Fatalf("Environment variable %s not defined", key)
	}

	return defValue
}
