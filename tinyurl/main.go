package main

import (
	"log"

	"tinyurl/helpers"
	"tinyurl/handlers"
	"tinyurl/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	uri := helpers.GetEnv()

	// connection function that returns the collection
	collection := helpers.ConnectToMongo(uri);

	// matching handlers to the mongo collections
	urlHandler := &handlers.URLHandler{Collection: collection}

	// initializing a gin router
	r := gin.Default()

	// registering the routes
	routes.URLRoute(r, urlHandler)

	log.Println("Server running on port 8080....")
	log.Fatal(r.Run(":8080"))
}