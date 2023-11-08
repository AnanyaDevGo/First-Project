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

	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, cfg, otpRepository, helper)
	userHandler := handler.NewUserHandler(userUseCase)

	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository, helper)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	inventoryRepository := repository.NewInventoryRepository(gormDB)
	inventoryUseCase := usecase.NewInventoryUseCase(inventoryRepository, helper)
	inventoryHandler := handler.NewInventoryHandler(inventoryUseCase)

	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, categoryHandler, otpHandler, inventoryHandler)

	return serverHTTP, nil
}
