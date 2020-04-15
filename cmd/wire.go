//+build wireinject

package main

import (
	projectHandler "todone/pkg/api/handler/project"
	userHandler "todone/pkg/api/handler/user"
	projectInteractor "todone/pkg/api/usecase/project"
	userInteractor "todone/pkg/api/usecase/user"
	"todone/pkg/domain/repository"
	projectSvc "todone/pkg/domain/service/project"
	userSvc "todone/pkg/domain/service/user"
	projectRepo "todone/pkg/infrastructure/mysql/project"
	userRepo "todone/pkg/infrastructure/mysql/user"

	"github.com/google/wire"
	"github.com/volatiletech/sqlboiler/boil"
)

func InitProjectAPI(db boil.ContextExecutor, masterTxManager repository.MasterTxManager) projectHandler.Server {
	wire.Build(projectRepo.New, projectSvc.New, projectInteractor.New, projectHandler.New)

	return projectHandler.Server{}
}

func InitUserAPI(db boil.ContextExecutor, masterTxManager repository.MasterTxManager) userHandler.Server {
	wire.Build(userRepo.New, userSvc.New, userInteractor.New, userHandler.New)

	return userHandler.Server{}
}
