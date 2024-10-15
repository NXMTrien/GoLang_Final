package routes

import (
	"restaurant-management/controllers"
	"restaurant-management/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", controllers.Register)
		userRoutes.POST("/login", controllers.Login)
		userRoutes.PUT("/update/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
	}
}
