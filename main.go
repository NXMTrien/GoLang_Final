package main

import (
	"restaurant-management/database"
	"restaurant-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()

	routes.UserRoutes(r)
	routes.FoodRoutes(r)
	routes.OrderRoutes(r)
	routes.InvoiceRoutes(r)
	// Thêm các nhóm route khác tại đây

	r.Run(":8181")
}
