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

// Store ...
type Store struct {
	Db     *mongo.Database
	Client *mongo.Client
}

// Connect - ...
func (s *Store) Connect(dbName string, connStr string) {
	var connectOnce sync.Once

	connectOnce.Do(func() {
		s.Db, s.Client = connectToMongo(dbName, connStr)
	})
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
