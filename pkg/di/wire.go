//go:build wireinject
// +build wireinject

package di

import (
	http "CrocsClub/pkg/api"
	"CrocsClub/pkg/api/handler"
	config "CrocsClub/pkg/config"
	db "CrocsClub/pkg/db"
	repository "CrocsClub/pkg/repository"
	usecase "CrocsClub/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
