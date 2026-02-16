package main

import (
	"log"
	"inventory/db"
	"inventory/handlers"
	"inventory/routes"

	"github.com/gin-gonic/gin"
)

var Inventory int = 100

func main() {
	// db connection
	orderCollection := db.ConnectToMongo()

	// defining the handlers and collections
	orderHandler := &handlers.OrderCollection {
		Collection: orderCollection,
	}

	r := gin.Default()

	routes.OrderRoute(r, orderHandler)

	log.Fatal(r.Run())
}
