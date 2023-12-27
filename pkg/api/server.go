package http

import (
	"CrocsClub/pkg/api/handler"
	"CrocsClub/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	categoryHandler *handler.CategoryHandler,
	otpHandler *handler.OtpHandler,
	inventoryHandler *handler.InventoryHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymentHandler,
	walletHandler *handler.WalletHandler,
	couponHandler *handler.CouponHandler,
	offerHandler *handler.OfferHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	// engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.LoadHTMLGlob("template/*")
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/validate-token", adminHandler.ValidateRefreshTokenAndCreateNewAccess)

	routes.UserRoutes(engine.Group("/user"), userHandler, otpHandler, inventoryHandler, cartHandler, orderHandler, paymentHandler, walletHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryHandler, inventoryHandler, orderHandler, couponHandler, offerHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":8080")

	if err != nil {
		log.Fatal("gin engine couldn't start")
	}
}
