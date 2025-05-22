//go:build wireinject
// +build wireinject

package main

import (
	"wire-gorm-server/db"
	"wire-gorm-server/handler"
	"wire-gorm-server/repository"
	"wire-gorm-server/service"

	"github.com/google/wire"
)

func InitUserHandler() *handler.UserHandler {
	wire.Build(
		db.ConnectGORM,
		repository.NewUserRepository,
		service.NewUserService,
		handler.NewUserHandler,
	)
	return nil
}
