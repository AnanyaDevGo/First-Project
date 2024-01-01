package routes

import (
	"CrocsClub/pkg/api/handler"
	"CrocsClub/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler, offerHandler *handler.OfferHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)
	engine.GET("/inventories", inventoryHandler.ListProducts)
	engine.Use(middleware.AdminAuthMiddleware)
	{
		engine.GET("/dashboard", adminHandler.Dashboard)

		engine.GET("/salesbydate", adminHandler.SalesByDate)
		engine.GET("/salesreport", adminHandler.FilteredSalesReport)
		engine.GET("/sales-report-date", adminHandler.SalesReportByDate)

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

			inventorymanagement.PUT("", inventoryHandler.EditInventory)
			inventorymanagement.DELETE("", inventoryHandler.DeleteInventory)
			inventorymanagement.PUT("/stock", inventoryHandler.UpdateInventory)
			inventorymanagement.POST("/uploadimages", inventoryHandler.MultipleImageUploader)
		}
		order := engine.Group("/order")
		{
			order.GET("/get", orderHandler.GetAdminOrders)
			order.POST("/status", orderHandler.ApproveOrder)

		}
		payment := engine.Group("/payment-method")
		{
			payment.POST("/pay", adminHandler.NewPaymentMethod)
			payment.GET("", adminHandler.ListPaymentMethods)
			payment.DELETE("", adminHandler.DeletePaymentMethod)
		}
		coupon := engine.Group("/coupon")
		{
			coupon.POST("", couponHandler.AddCoupon)
			coupon.GET("", couponHandler.GetCoupon)
			coupon.PATCH("", couponHandler.EditCoupon)
		}
		offer := engine.Group("offer")
		{
			offer.POST("/product-offer", offerHandler.AddProductOffer)
			offer.GET("/get-product-offer", offerHandler.GetProductOffer)
			offer.DELETE("/expire-product-offer", offerHandler.ExpireProductOffer)

			offer.POST("/category-offer", offerHandler.AddCategoryOffer)
			offer.GET("/get-category-offer", offerHandler.GetCategoryOffer)
			offer.DELETE("/expire-category-offer", offerHandler.ExpireCategoryOffer)
		}

	}
}
