package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserName    string             `json:"userName" bson:"userName"`
	Address     string             `json:"address" bson:"address"`
	ProductName string             `json:"productName" bson:"productName"`
}
