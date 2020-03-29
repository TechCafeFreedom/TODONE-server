//+build wireinject

package main

import (
	userHandler "todone/pkg/api/handler/user"
	userInteractor "todone/pkg/api/usecase/user"
	userSvc "todone/pkg/domain/service/user"
	userRepo "todone/pkg/infrastructure/mysql/user"

	"github.com/google/wire"
	"github.com/volatiletech/sqlboiler/boil"
)

func InitUserAPI(db boil.ContextExecutor) userHandler.Server {
	wire.Build(userRepo.New, userSvc.New, userInteractor.New, userHandler.New)

	return userHandler.Server{}
}
