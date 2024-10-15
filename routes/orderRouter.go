package routes

import (
	"restaurant-management/controllers"
	// "restaurant-management/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	// Order Routes
	orderRoutes := router.Group("/orders")
	orderRoutes.GET("/", controllers.GetAllOrders)
	orderRoutes.GET("/:id", controllers.GetOrderByID) // Thêm dòng này
	orderRoutes.POST("/", controllers.CreateOrder)
	orderRoutes.PUT("/:id", controllers.UpdateOrder)    // Thêm dòng này
	orderRoutes.DELETE("/:id", controllers.DeleteOrder) // Thêm dòng này

}
