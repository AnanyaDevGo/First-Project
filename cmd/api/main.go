package main

import (
	"CrocsClub/pkg/config"
	"CrocsClub/pkg/di"
	"log"

	_ "CrocsClub/cmd/api/docs"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title Go + Gin E-Commerce API Crocs Club
// @version 1.0.0
// @description Crocs Club is an E-commerce platform to purchase Crocs
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)

	if diErr != nil {

		log.Fatal("cannot start server: ", diErr)
	} else {
		//server.SetupSwagger()

		server.Start()
	}
}
