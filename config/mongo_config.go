package config

import (
	"context"
	"log"
	"os"

	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_URI      = os.Getenv("MONGO_URI")
	MONGO_DATABASE = os.Getenv("MONGO_DATABASE")
)

type MongoConfig struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func (m *MongoConfig) Start() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalf("Failed to connect to mongo")
	}

	m.Client = client
	m.Database = m.Client.Database(MONGO_DATABASE)
}

func (m *MongoConfig) Close() {
	m.Client.Disconnect(context.Background())
}
