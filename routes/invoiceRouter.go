package routes

import (
	"restaurant-management/controllers"
	// "restaurant-management/middlewares"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(router *gin.Engine) {
	invoiceRoutes := router.Group("/invoices")
	invoiceRoutes.GET("/", controllers.GetAllInvoices)
	invoiceRoutes.GET("/:id", controllers.GetInvoiceByID)
	invoiceRoutes.POST("/", controllers.CreateInvoice)
	invoiceRoutes.PUT("/:id", controllers.UpdateInvoice)
	invoiceRoutes.DELETE("/:id", controllers.DeleteInvoice)
}
