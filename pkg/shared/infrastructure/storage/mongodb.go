package storage

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	timeOutDuration = 5 * time.Second
)

var mongoOnce sync.Once
var clientInstance *mongo.Client

// InitializeMongoConnection Initialize connection
func InitializeMongoConnection() {
	log.Println("Starting mongo connection...")
	client := NewMongoClient()

	if client == nil {
		log.Println("Mongo client initialized")
		return
	}
}

// NewMongoClient Initialize mongo client
func NewMongoClient() *mongo.Client {
	mongoOnce.Do(func() {
		stringConnection := getStringConnection()
		client, err := mongo.NewClient(options.Client().ApplyURI(stringConnection))

		if err != nil {
			panic(err)
		}

		ctx := getNewContext()
		err = client.Connect(ctx)

		if err != nil {
			panic(err)
		}

		err = client.Ping(context.Background(), readpref.Primary())

		if err != nil {
			panic(err)
		}

		clientInstance = client
	})

	return clientInstance
}

func getStringConnection() string {
	return "mongodb://" + os.Getenv("MONGO_USERNAME") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST")
}

func getNewContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeOutDuration)

	defer cancel()

	return ctx
}
