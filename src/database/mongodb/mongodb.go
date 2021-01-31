package mongodb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect - ...
func Connect(dbName string, connStr string) (*mongo.Database, *mongo.Client) {
	var (
		connectOnce sync.Once
		db          *mongo.Database
		client      *mongo.Client
	)

	connectOnce.Do(func() {
		db, client = connectToMongo(dbName, connStr)
	})

	if db == nil || client == nil {
		logrus.Fatal("Failed connect to database")
	}

	return db, client
}

func connectToMongo(dbName string, connStr string) (*mongo.Database, *mongo.Client) {
	var err error

	client, err := mongo.NewClient(options.Client().ApplyURI(connStr))
	if err != nil {
		logrus.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logrus.Fatal(err)
	}

	var db = client.Database(dbName)
	fmt.Printf("Connected to DbName: %s\n", dbName)

	return db, client
}
