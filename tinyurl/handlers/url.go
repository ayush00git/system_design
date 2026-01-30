package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"tinyurl/internals/helpers"
	"tinyurl/internals/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLHandler struct {
	Collection *mongo.Collection
}

func (h *URLHandler) ToTinyURL(w http.ResponseWriter, r *http.Request) {

	// url data structure which handles url type struct
	var url models.URL

	// decoding the json response sent by user
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, "Error decoding user's request", http.StatusInternalServerError)
		return
	}

	// defining the mongo fields
	url.ID = primitive.NewObjectID()
	url.CreatedAt = time.Now()

	// get the 'LongURL' and generate a 'TinyURL' from it

	// send it to mongodb
	_, err := h.Collection.InsertOne(context.TODO(), url)
	if err != nil {
		http.Error(w, "Error inserting the document to mongodb", http.StatusInternalServerError)
		return
	}
	// return some response
	response := map[string]interface{} {
		"message": "tiny url generated successfully!",
		"url": url,
	}

	
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding server's response", http.StatusInternalServerError)
		return
	}
}

func (h *URLHandler) GenerateTinyURL (c *gin.Context) {
	var url models.URL

	// decoding the user's request body
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid response",
		})
		return
	}

	url.ID = primitive.NewObjectID()
	url.ShortCode = helpers.GenerateCode()
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

func (h *URLHandler) HitTinyURL (w http.ResponseWriter, r *http.Request) {

}