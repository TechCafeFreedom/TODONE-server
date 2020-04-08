//+build wireinject

package main

import (
	projectHandler "todone/pkg/api/handler/project"
	userHandler "todone/pkg/api/handler/user"
	projectInteractor "todone/pkg/api/usecase/project"
	userInteractor "todone/pkg/api/usecase/user"
	projectSvc "todone/pkg/domain/service/project"
	userSvc "todone/pkg/domain/service/user"
	"todone/pkg/infrastructure/mysql"
	projectRepo "todone/pkg/infrastructure/mysql/project"
	userRepo "todone/pkg/infrastructure/mysql/user"

	"github.com/google/wire"
	"github.com/volatiletech/sqlboiler/boil"
)

func InitProjectAPI(db boil.ContextExecutor, masterTxManager mysql.MasterTxManager) projectHandler.Server {
	wire.Build(projectRepo.New, projectSvc.New, projectInteractor.New, projectHandler.New)

	return projectHandler.Server{}
}

func InitUserAPI(db boil.ContextExecutor, masterTxManager mysql.MasterTxManager) userHandler.Server {
	wire.Build(userRepo.New, userSvc.New, userInteractor.New, userHandler.New)

	return userHandler.Server{}
}
