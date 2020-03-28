//+build wireinject

package main

import (
	"todone/pkg/api/handler"
	userSvc "todone/pkg/domain/service/user"
	userRepo "todone/pkg/infrastructure/mysql/user"

	"github.com/google/wire"
	"github.com/volatiletech/sqlboiler/boil"
)

func InitUserAPI(db boil.ContextExecutor) handler.UserHandler {
	wire.Build(userRepo.NewUserRepoImpl, userSvc.NewUserService, handler.NewUserHandler)

	return handler.UserHandler{}
}
