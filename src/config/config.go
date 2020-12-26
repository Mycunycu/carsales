package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	Port string
}

func Init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("No such file .env")
	}
}

func getEnv(envName string, defaultVal string) string{
	if env, exist := os.LookupEnv(envName); exist {
		return env
	}

	return defaultVal
}

func GetConfig() *Config {
	return &Config{
		Port: getEnv("PORT", "8000"),
	}
}
