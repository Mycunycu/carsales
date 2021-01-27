package config

import (
	"carsales/models"

	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Init ...
func Init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("No such file .env")
	}
}

func getEnv(envName string, defaultVal string) string {
	if env, exist := os.LookupEnv(envName); exist {
		return env
	}

	return defaultVal
}

// GetConfig ...
func GetConfig() *models.Config {
	return &models.Config{
		Port:       getEnv("PORT", "8000"),
		DBHost:     getEnv("DBHOST", ""),
		DBPort:     getEnv("DBPORT", ""),
		DBUserName: getEnv("DBUSERNAME", ""),
		DBPassword: getEnv("DBPASSWORD", ""),
		DBName:     getEnv("DBNAME", ""),
		SSLMode:    getEnv("SSLMODE", ""),
	}
}
