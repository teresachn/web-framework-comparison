package config

import (
	"context"
	"log"
	"os"

	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGO_URI = os.Getenv("MONGO_URI")
)

type Configuration struct {
	Client *mongo.Client
}

func (m *Configuration) StartDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MONGO_URI))
	if err != nil {
		log.Fatalf("Failed to connect to mongo")
	}

	m.Client = client
}

func (m *Configuration) CloseDB() {
	m.Client.Disconnect(context.Background())
}
