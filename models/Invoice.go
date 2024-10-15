package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Invoice struct represents an invoice in the database
type Invoice struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	TotalAmount float64            `bson:"total_amount" json:"total_amount"` // Thêm trường TotalAmount
	Status      string             `bson:"status" json:"status"`             // Thêm trường Status
	CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at" json:"updated_at"`
}
