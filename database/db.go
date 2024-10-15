package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Kết nối tới MongoDB
func ConnectDB() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB: ", err)
	} else {
		log.Println("MongoDB client created successfully")
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed: ", err)
	} else {
		log.Println("MongoDB ping successful")
	}

	log.Println("Connected to MongoDB!")
	Client = client
}
func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("Golang_Final").Collection(collectionName)
}
