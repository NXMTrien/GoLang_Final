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

// GetAllFoods retrieves all food items
func GetAllFoods(c *gin.Context) {
	var foods []models.Food
	collection := database.GetCollection("foods")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var food models.Food
		if err := cursor.Decode(&food); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		foods = append(foods, food)
	}

	c.JSON(http.StatusOK, foods)
}

// GetFoodByID retrieves a single food item by ID
func GetFoodByID(c *gin.Context) {
	foodID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(foodID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
		return
	}

	var food models.Food
	collection := database.GetCollection("foods")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&food)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Food not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, food)
}

// CreateFood creates a new food item
func CreateFood(c *gin.Context) {
	var food models.Food
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.GetCollection("foods")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, food)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"foodId": result.InsertedID})
}

// UpdateFood updates a food item by ID
func UpdateFood(c *gin.Context) {
	foodID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(foodID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
		return
	}

	var foodUpdate models.Food
	if err := c.ShouldBindJSON(&foodUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.GetCollection("foods")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":        foodUpdate.Name,
			"price":       foodUpdate.Price,
			"description": foodUpdate.Description,
			"image":       foodUpdate.Image,
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food updated successfully"})
}

// DeleteFood deletes a food item by ID
func DeleteFood(c *gin.Context) {
	foodID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(foodID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid food ID"})
		return
	}

	collection := database.GetCollection("foods")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food deleted successfully"})
}
