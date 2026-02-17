package handlers

import (
	"net/http"
	"context"

	"inventory/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	Collection *mongo.Collection
}

func (h *InventoryHandler) PostAProduct (c *gin.Context) {
	var inv models.Inventory
	if err := c.ShouldBindJSON(&inv); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request body"})
		return
	}

	inv.ID = primitive.NewObjectID()
	inv.Units = 100							// hardcoded for now
	inv.ProductName = "Apple IPhone 16"		// hardcoded for now

	_, err := h.Collection.InsertOne(context.TODO(), inv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting the document"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Inventory updated with a new item!",
		"product": inv,
	})
}
