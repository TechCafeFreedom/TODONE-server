// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/volatiletech/sqlboiler/boil"
	"todone/pkg/api/handler/user"
	user4 "todone/pkg/api/usecase/user"
	user3 "todone/pkg/domain/service/user"
	user2 "todone/pkg/infrastructure/mysql/user"
)

// Injectors from wire.go:

func InitUserAPI(db boil.ContextExecutor) user.Server {
	repository := user2.New(db)
	service := user3.New(repository)
	interactor := user4.New(service)
	server := user.New(interactor)
	return server
}