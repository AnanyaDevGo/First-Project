package routes

import (
	"CrocsClub/pkg/api/handler"
	"CrocsClub/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		engine.GET("/dashboard", adminHandler.Dashboard)
		engine.GET("/salesreport", adminHandler.FilteredSalesReport)

		usermanagement := engine.Group("/users")
		{
			usermanagement.GET("/", adminHandler.GetUsers)
			usermanagement.PUT("/block", adminHandler.BlockUser)
			usermanagement.PUT("/unblock", adminHandler.UnBlockUser)
		}

		categorymanagement := engine.Group("/category")
		{
			categorymanagement.GET("", categoryHandler.GetCategory)
			categorymanagement.POST("", categoryHandler.AddCategory)
			categorymanagement.PUT("", categoryHandler.UpdateCategory)
			categorymanagement.DELETE("", categoryHandler.DeleteCategory)

		}
		inventorymanagement := engine.Group("/inventories")
		{
			inventorymanagement.POST("", inventoryHandler.AddInventory)
			inventorymanagement.DELETE("", inventoryHandler.DeleteInventory)
			//inventorymanagement.GET("", inventoryHandler.ListProductsForAdmin)

			inventorymanagement.PUT("/stock", inventoryHandler.UpdateInventory)
		}
		order := engine.Group("/order")
		{
			order.GET("/get", orderHandler.GetAdminOrders)
			order.GET("/status", orderHandler.ApproveOrder)

		}
		payment := engine.Group("/payment-method")
		{
			payment.POST("/pay", adminHandler.NewPaymentMethod)
			payment.GET("", adminHandler.ListPaymentMethods)
			payment.DELETE("", adminHandler.DeletePaymentMethod)
		}

	}
}
