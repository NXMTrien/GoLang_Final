package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `json:"name"`
	Image       string             `json:"image"`
	Price       float64            `json:"price"`
	Description string             `json:"description"`
}
