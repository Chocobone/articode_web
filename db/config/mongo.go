package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var UserCollection *mongo.Collection

func InitMongo() {
	uri := "mongodb://localhost:27017" //os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("ERROR: 'MONGO_URI' is NOT set.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("ERROR: failed to connect to MongoDB: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("ERROR: failed to ping MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB Atlas")

	Client = client

	//initialize collection
	UserCollection = Client.Database("resq").Collection("users")
	//add collection if struct added
}
