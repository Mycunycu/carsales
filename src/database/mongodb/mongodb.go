package mongodb

import (
	"carsales/internal/config"
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Store ...
type MongoStore struct {
	*mongo.Database
}

// Connect - ...
func Connect(cfg *config.Config) (*MongoStore, error) {
	var once sync.Once
	var db *mongo.Database
	var err error

	once.Do(func() {
		db, err = connectToMongo(cfg.MongoConnStr, cfg.DbName)
	})

	if err != nil {
		return nil, err
	}

	return &MongoStore{db}, nil
}

func connectToMongo(connStr string, dbName string) (*mongo.Database, error) {
	var err error

	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	var db = client.Database(dbName)
	fmt.Printf("Connected to DbName: %s\n", dbName)

	return db, nil
}
