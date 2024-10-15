package controllers

import (
	"context"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"restaurant-management/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Register handles user registration
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword
	user.Role = "001" // Đặt vai trò mặc định

	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"userId": result.InsertedID})
}

// Login handles user login
func Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"email": loginData.Email}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, _ := utils.GenerateToken(user.ID.Hex())
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// UpdateUser updates user information
func UpdateUser(c *gin.Context) {
	var userUpdate models.User
	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Param("id")
	oid, _ := primitive.ObjectIDFromHex(userID)

	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":  userUpdate.Name,
			"email": userUpdate.Email,
		},
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cập nhật người dùng thành công"})
}
