package db

import (
	"context"
	"time"
	"os"
	"log"
	"fmt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoURI() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading the .env file")
	}

	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Unable to find MONGO_URI")
	}
	return uri
}

func ConnectToMongo() *mongo.Database {
	// ctx, client, ping, return the collection

	uri := GetMongoURI()
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("Connection timeout error: %s", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Connection to mongodb is failing")
	}
	fmt.Println("Connected to MongoDB!")

	database := client.Database("test")
	return database
}
