package handlers

import (
	"net/http"
	"context"

	"inventory/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Collection *mongo.Collection
}

func (h *OrderHandler) HealthRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "All good bruhh!"})
}

func (h *OrderHandler) PlaceOrder (c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
		return
	}
	
	filter := bson.M{
		"productName": "Apple IPhone 16",
		"units": bson.M{"$gt": 0},
	}

	update := bson.M{
		"$inc": bson.M{"units": -1},
	}
	// give the key to a specific goroutine
	var updatedInventory models.Inventory
	err := h.Collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&updatedInventory)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "sold out!"})
		return
	}

	leftUnits := updatedInventory.Units

	order.ID = primitive.NewObjectID()
	order.ProductName = "Apple IPhone 16"

	_, err = h.Collection.InsertOne(context.TODO(), order)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "sold out!"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Order placed!", 
		"left_units": leftUnits,
	})
}
