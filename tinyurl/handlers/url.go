package handlers

import (
	"context"
	"net/http"
	"time"

	"tinyurl/helpers"
	"tinyurl/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type URLHandler struct {
	Collection *mongo.Collection
}

func (h *URLHandler) ToTinyURL (c *gin.Context) {
	var url models.URL

	// decoding the user's request body
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid response",
		})
		return
	}

	url.ID = primitive.NewObjectID()
	url.TinyId = helpers.GenerateCode(6)
	url.TinyURL = "http://localhost:8080/" + url.TinyId
	url.CreatedAt = time.Now()

	// save request to the database
	_, err := h.Collection.InsertOne(context.TODO(), url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error saving to the database",
		})
		return
	}

	// send in response
	c.JSON(http.StatusOK, gin.H{
		"message": "Tiny URL generated!",
		"url": url,
	})
}

func (h *URLHandler) HitTinyURL (c *gin.Context) {
	tinyId := c.Param("tinyId")

	var foundURL models.URL

	filter := bson.M{"tinyId": tinyId}
	err := h.Collection.FindOne(context.TODO(), filter).Decode(&foundURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Finding the redirection url",
		})
	}

	// http.StatusFound or status code 302 is a temporary redirect used for keeping analytics,
	// status code 301 can be used for a permanent redirection and would be fast, but we'll
	// lost a track of clicks
	c.Redirect(http.StatusFound, foundURL.LongURL)	
}
