package main

import (
	"log"

	"tinyurl/helpers"
	"tinyurl/handlers"
	"tinyurl/routes"
	"tinyurl/models"

	"github.com/gin-gonic/gin"
)

func main() {
	uri := helpers.GetEnv()

	// connection function that returns the collection
	collection := helpers.ConnectToMongo(uri);

	// creating a buffer of 5000, means the server will accept 
	// 5000 requests even if the mongodb is not connected
	queue := make(chan models.URL, 15000)

	// matching handlers to the mongo collections
	urlHandler := &handlers.URLHandler{
		Collection: collection,
		Queue: queue,
	}

	// this spawns the 50 concurrent goroutine that lives forever
	urlHandler.StartWorker(50)

	// initializing a gin router
	r := gin.Default()

	// registering the routes
	routes.URLRoute(r, urlHandler)

	log.Println("Server running on port 8080....")
	log.Fatal(r.Run(":8080"))
}