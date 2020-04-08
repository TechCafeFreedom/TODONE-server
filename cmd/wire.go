//+build wireinject

package main

import (
	projectHandler "todone/pkg/api/handler/project"
	projectInteractor "todone/pkg/api/usecase/project"
	projectSvc "todone/pkg/domain/service/project"
	"todone/pkg/infrastructure/mysql"
	projectRepo "todone/pkg/infrastructure/mysql/project"

	"github.com/google/wire"
	"github.com/volatiletech/sqlboiler/boil"
)

func InitProjectAPI(db boil.ContextExecutor, masterTxManager mysql.MasterTxManager) projectHandler.Server {
	wire.Build(projectRepo.New, projectSvc.New, projectInteractor.New, projectHandler.New)

	return projectHandler.Server{}
}
