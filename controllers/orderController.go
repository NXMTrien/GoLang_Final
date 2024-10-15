package controllers

import (
	"context"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllOrders retrieves all orders for the authenticated user
func GetAllOrders(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var orders []models.Order
	collection := database.GetCollection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Chuyển đổi userID sang primitive.ObjectID
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	cursor, err := collection.Find(ctx, bson.M{"user_id": oid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		orders = append(orders, order)
	}

	c.JSON(http.StatusOK, orders)
}

// CreateOrder creates a new order for the authenticated user
func CreateOrder(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chuyển đổi userID sang primitive.ObjectID
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	order.UserID = oid
	collection := database.GetCollection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"orderId": result.InsertedID})
}

// GetOrderByID retrieves a single order by ID
func GetOrderByID(c *gin.Context) {
	orderID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	collection := database.GetCollection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

// UpdateOrder updates an order by ID
func UpdateOrder(c *gin.Context) {
	orderID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var orderUpdate models.Order
	if err := c.ShouldBindJSON(&orderUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.GetCollection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": orderUpdate,
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// DeleteOrder deletes an order by ID
func DeleteOrder(c *gin.Context) {
	orderID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	collection := database.GetCollection("orders")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
