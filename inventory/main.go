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

	orderCollec := database.Collection("orders")
	inventoryCollec := database.Collection("inventory")

	// defining the handlers and collections
	orderHandler := &handlers.OrderHandler{
		OrderCollection: orderCollec,
		InventoryCollection: inventoryCollec,
	}

	// inventoryHandler := &handlers.InventoryHandler {
	// 	Collection: inventoryCollection,
	// }

	r := gin.Default()

	routes.OrderRoute(r, orderHandler)
	// routes.InventoryRoutes(r, inventoryHandler)

	log.Fatal(r.Run())
}
