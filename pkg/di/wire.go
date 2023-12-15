//go:build wireinject
// +build wireinject

package di

import (
	http "CrocsClub/pkg/api"
	"CrocsClub/pkg/api/handler"
	config "CrocsClub/pkg/config"
	db "CrocsClub/pkg/db"
	"CrocsClub/pkg/helper"
	repository "CrocsClub/pkg/repository"
	usecase "CrocsClub/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,

		repository.NewUserRepository,
		repository.NewAdminRepository,
		repository.NewOtpRepository,
		repository.NewCategoryRepository,
		repository.NewInventoryRepository,
		repository.NewOrderRepository,
		repository.NewCartRepository,
		repository.NewPaymentRepository,
		repository.NewWalletRepository,
		repository.NewCouponRepository,
		repository.NewOfferRepository,

		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		usecase.NewOtpUseCase,
		usecase.NewCategoryUseCase,
		usecase.NewInventoryUseCase,
		usecase.NewOrderUseCase,
		usecase.NewCartUseCase,
		usecase.NewPaymentUseCase,
		usecase.NewWalletUseCase,
		usecase.NewCouponUsecase,
		usecase.NewOferUsecase,

		handler.NewUserHandler,
		handler.NewAdminHandler,
		handler.NewOtpHandler,
		handler.NewCategoryHandler,
		handler.NewInventoryHandler,
		handler.NewOrderHandler,
		handler.NewCartHandler,
		handler.NewPaymentHandler,
		handler.NewWalletHandler,
		handler.NewCouponHandler,
		handler.NewOfferHandler,

		helper.NewHelper,

		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
