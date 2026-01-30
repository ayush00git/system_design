package models

import (
	"time"
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type URL struct {
	ID				primitive.ObjectID			`json:"_id" bson:"_id, omitempty"`
	LongURL			string						`json:"longUrl" bson:"longUrl" binding:"required"`
	ShortCode		string						`json:"shortCode" bson:"shortCode"`
	TinyURL			string						`json:"tinyUrl" bson:"tinyUrl"`
	CreatedAt		time.Time					`json:"created_at" bson:"created_at"`
}
