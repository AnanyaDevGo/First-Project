package di

import (
	http "CrocsClub/pkg/api"
	"CrocsClub/pkg/api/handler"
	"CrocsClub/pkg/config"
	"CrocsClub/pkg/db"
	"CrocsClub/pkg/repository"
	"CrocsClub/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUseCase(cfg, otpRepository)
	otpHandler := handler.NewOtpHandler(otpUseCase)

	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, cfg, otpRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	serverHTTP := http.NewServerHTTP(userHandler, adminHandler, categoryHandler, otpHandler)

	return serverHTTP, nil
}
