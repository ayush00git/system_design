package main

import (
	"log"
	"inventory/db"
	"inventory/handlers"
	"inventory/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// db connection
	database := db.ConnectToMongo()

	orderCollection := database.Collection("orders")
	inventoryCollection := database.Collection("inventory")

	// defining the handlers and collections
	orderHandler := &handlers.OrderHandler {
		Collection: orderCollection,
	}

	inventoryHandler := &handlers.InventoryHandler{
		Collection: inventoryCollection,
	}

	r := gin.Default()

	routes.OrderRoute(r, orderHandler)
	routes.InventoryRoutes(r, inventoryHandler)

	log.Fatal(r.Run())
}
