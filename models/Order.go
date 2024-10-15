package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID   primitive.ObjectID `json:"userId"`
	FoodID   primitive.ObjectID `json:"foodId"`
	Quantity int                `json:"quantity"`
	Status   string             `json:"status"` // e.g., "pending", "paid"
}
