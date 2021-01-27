package database

import (
	"database/sql"
	"fmt"

	"carsales/models"

	_ "github.com/lib/pq" // ....
	"github.com/sirupsen/logrus"
)

// Connect ...
func Connect(cfg *models.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUserName, cfg.DBName, cfg.DBPassword, cfg.SSLMode)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logrus.Println("Connection to DB succesfully")
	return db, nil
}
