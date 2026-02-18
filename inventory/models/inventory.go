package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Inventory struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	ProductName string             `json:"productName" bson:"productName"`
	Units       int                `json:"units" bson:"units"`
}
