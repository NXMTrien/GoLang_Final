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

// GetAllInvoices retrieves all invoices for the authenticated user
func GetAllInvoices(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var invoices []models.Invoice
	collection := database.GetCollection("invoices")
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
		var invoice models.Invoice
		if err := cursor.Decode(&invoice); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		invoices = append(invoices, invoice)
	}

	c.JSON(http.StatusOK, invoices)
}

// CreateInvoice creates a new invoice for the authenticated user
func CreateInvoice(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chuyển đổi userID sang primitive.ObjectID
	oid, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	invoice.UserID = oid
	collection := database.GetCollection("invoices")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, invoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"invoiceId": result.InsertedID})
}

// GetInvoiceByID retrieves a single invoice by ID
func GetInvoiceByID(c *gin.Context) {
	invoiceID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(invoiceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	var invoice models.Invoice
	collection := database.GetCollection("invoices")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&invoice)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

// UpdateInvoice updates an invoice by ID
func UpdateInvoice(c *gin.Context) {
	invoiceID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(invoiceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	var invoiceUpdate models.Invoice
	if err := c.ShouldBindJSON(&invoiceUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := database.GetCollection("invoices")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"total_amount": invoiceUpdate.TotalAmount,
			"status":       invoiceUpdate.Status,
			// Thêm các trường khác nếu cần
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": oid}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice updated successfully"})
}

// DeleteInvoice deletes an invoice by ID
func DeleteInvoice(c *gin.Context) {
	invoiceID := c.Param("id")
	oid, err := primitive.ObjectIDFromHex(invoiceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	collection := database.GetCollection("invoices")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice deleted successfully"})
}
