package helpers

import (
	"fmt"
	"os"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

func GetEnv() (string) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error while loading the .env: %s", err)
		os.Exit(1)
	}

	uri := os.Getenv("MONGO_URI");
	if uri == "" {
		fmt.Println("Unable to find MONGO_URI in .env")
		os.Exit(1)
	}
	return uri
}

func ConnectToMongo (uri string) {
	// checking the connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// connecting to mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Printf("Error connecting to mongodb client: %s", err)
		os.Exit(1)
	}

	// checking for ping
	if err := client.Ping(ctx, nil); err != nil {
		fmt.Printf("Error making a connection to mongodb: %s", err)
		os.Exit(1)
	}
	fmt.Println("Connected to mongodb")
	
	// selecting the dbs and collections
}