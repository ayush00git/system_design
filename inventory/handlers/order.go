package handlers

import (
	"net/http"
	"context"

	"inventory/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gin-gonic/gin"
)

type OrderCollection struct {
	Collection *mongo.Collection
}

func (h *OrderCollection) HealthRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "All good bruhh!"})
}

func (h *OrderCollection) PlaceOrder (c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "decoding the response"})
	}

	order.ID = primitive.NewObjectID()
	order.ProductName = "Apple IPhone 16"

	unit, err := h.Collection.InsertOne(context.TODO(), order)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Error inserting to the db"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Order placed!", "order_details": unit})
}
