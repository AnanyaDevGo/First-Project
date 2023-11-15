package routes

import (
	"CrocsClub/pkg/api/handler"
	"CrocsClub/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, otpHandler *handler.OtpHandler, inventoryHandler *handler.InventoryHandler, cartHandler *handler.CartHandler, orderHandler *handler.OrderHandler) {

	engine.POST("/signup", userHandler.UserSignUp)
	engine.POST("/login", userHandler.LoginHandler)

	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/verifyotp", otpHandler.VerifyOTP)

	engine.Use(middleware.UserAuthMiddleware)

	profile := engine.Group("/profile")
	{
		profile.GET("/address", userHandler.GetAddress)
		profile.POST("/addaddress", userHandler.AddAddress)
		profile.GET("/details", userHandler.GetUserDetails)
	}
	edit := engine.Group("/edit")
	{
		edit.PUT("/name", userHandler.EditName)
		edit.PUT("/email", userHandler.EditEmail)
		edit.PUT("/phone", userHandler.EditPhone)
	}

	cart := engine.Group("/cart")
	{
		cart.GET("/", userHandler.GetCart)
		cart.DELETE("/remove", userHandler.RemoveFromCart)
		cart.PUT("/updateQuantity/plus", userHandler.UpdateQuantityAdd)
		cart.PUT("/updateQuantity/minus", userHandler.UpdateQuantityLess)

	}

	checkout := engine.Group("/check-out")
	{
		checkout.GET("", cartHandler.CheckOut)
	}
	order := engine.Group("/order")
	{
		order.POST("", orderHandler.GetOrders)
		order.DELETE("", orderHandler.CancelOrder)
	}

}
