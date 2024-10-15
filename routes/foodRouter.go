package routes

import (
	"restaurant-management/controllers"
	// "restaurant-management/middlewares"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
	foodRoutes := router.Group("/foods")
	foodRoutes.GET("/", controllers.GetAllFoods)
	foodRoutes.GET("/:id", controllers.GetFoodByID)
	foodRoutes.POST("/", controllers.CreateFood)
	foodRoutes.PUT("/:id", controllers.UpdateFood)
	foodRoutes.DELETE("/:id", controllers.DeleteFood)
}
