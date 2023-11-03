package http

import (
	"CrocsClub/pkg/api/handler"
	"CrocsClub/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	 adminHandler *handler.AdminHandler,
	 categoryHandler *handler.CategoryHandler,
	  otpHandler *handler.OtpHandler) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	routes.UserRoutes(engine.Group("/user"), userHandler, otpHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	err := sh.engine.Run(":3000")
	
	if err != nil {
		log.Fatal("gin engine couldn't start")
	}
}
