package di

import (
	http "CrocsClub/pkg/api"
	"CrocsClub/pkg/api/handler"
	"CrocsClub/pkg/config"
	"CrocsClub/pkg/db"
	"CrocsClub/pkg/helper"
	"CrocsClub/pkg/repository"
	"CrocsClub/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	helper := helper.NewHelper(cfg)

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, helper)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUseCase(cfg, otpRepository, helper)
	otpHandler := handler.NewOtpHandler(otpUseCase)

	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository, helper)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	inventoryRepository := repository.NewInventoryRepository(gormDB)
	inventoryUseCase := usecase.NewInventoryUseCase(inventoryRepository, helper)
	inventoryHandler := handler.NewInventoryHandler(inventoryUseCase)

	orderRepository := repository.NewOrderRepository(gormDB)

	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, cfg, otpRepository, inventoryRepository, orderRepository, helper)
	userHandler := handler.NewUserHandler(userUseCase)

	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository, inventoryRepository, userUseCase)
	cartHandler := handler.NewCartHandler(cartUseCase)

	orderUseCase := usecase.NewOrderUseCase(orderRepository, userUseCase)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, categoryHandler, otpHandler, inventoryHandler, cartHandler, orderHandler)

	return serverHTTP, nil
}
