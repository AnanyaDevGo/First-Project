package routes

import (
	"CrocsClub/pkg/api/handler"
	"CrocsClub/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, inventoryHandler *handler.InventoryHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler, walletHandler *handler.WalletHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.LoginHandler)

	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)

	engine.GET("/razorpay", paymentHandler.MakePaymentRazorPay)
	engine.GET("/status_update", paymentHandler.VerifyPayment)
	engine.GET("/home/list", inventoryHandler.ListProducts)

	engine.Use(middleware.UserAuthMiddleware)

	profile := engine.Group("/profile")
	{
		profile.GET("/address", userHandler.GetAddress)
		profile.POST("/addaddress", userHandler.AddAddress)
		profile.GET("/details", userHandler.GetUserDetails)

		order := profile.Group("/order")
		{
			order.GET("/get", orderHandler.GetOrders)
			order.GET("/all", orderHandler.GetAllOrders)
			order.DELETE("", orderHandler.CancelOrder)
			order.PATCH("/return", orderHandler.ReturnOrder)
		}
		edit := profile.Group("/edit")
		{
			edit.PATCH("/", userHandler.Edit)
		}
		security := profile.Group("/security")
		{
			security.PUT("/change-password", userHandler.ChangePassword)
		}
	}

	home := engine.Group("/home")
	{
		home.POST("/addcart", cartHandler.AddToCart)
		
	}
	cart := engine.Group("/cart")
	{
		cart.GET("/", userHandler.GetCart)
		cart.DELETE("/remove", userHandler.RemoveFromCart)
		cart.PUT("", userHandler.UpdateQuantity)
	}
	checkout := engine.Group("/check-out")
	{
		checkout.GET("", cartHandler.CheckOut)
		checkout.POST("/order", orderHandler.OrderItemsFromCart)
		checkout.GET("/print", orderHandler.PrintInvoice)
	}
	wallet := engine.Group("/wallet")
	{
		wallet.GET("", walletHandler.GetWallet)
		wallet.GET("/history", walletHandler.WalletHistory)
	}
	product := engine.Group("/product")
	{
		product.GET("/filter", inventoryHandler.FilterCategory)
		product.POST("search", inventoryHandler.SearchProducts)
	}

}
