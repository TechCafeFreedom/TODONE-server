//+build wireinject

package main

import (
	"todone/pkg/api/handler/user"
	userSvc "todone/pkg/domain/service/user"
	userRepo "todone/pkg/infrastructure/mysql/user"

	"github.com/google/wire"
	"github.com/volatiletech/sqlboiler/boil"
)

func InitUserAPI(db boil.ContextExecutor) user.Server {
	wire.Build(userRepo.NewUserRepoImpl, userSvc.NewUserService, user.NewUserHandler)

	return user.Server{}
}
