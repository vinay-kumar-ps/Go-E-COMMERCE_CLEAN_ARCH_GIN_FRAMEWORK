//go:build wireinject
// +build wireinject


package di

import (
	http "ecommerce/pkg/api"
	handler"ecommerce/pkg/api/handler"
    config	"ecommerce/pkg/config"
    db	"ecommerce/pkg/db"
    repository	"ecommerce/pkg/repository"
	usecase "ecommerce/pkg/usecase"

	"github.com/google/wire"
)


func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}