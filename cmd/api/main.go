package main

import (
	"CrocsClub/pkg/config"
	"CrocsClub/pkg/di"

	"log"

	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// @title Go + Gin E-Commerce API
// @version 1.0.0
// @description TechDeck is an E-commerce platform to purchase and sell Electronic itmes
// @contact.name API Support
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	// // swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "CrocsClub - E-commerce"
	docs.SwaggerInfo.Description = "CrocsClub is an E-commerce platform to purchasing and selling crocs"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
