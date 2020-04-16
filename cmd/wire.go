//+build wireinject

package main

import (
	boardHandler "todone/pkg/api/handler/board"
	userHandler "todone/pkg/api/handler/user"
	boardInteractor "todone/pkg/api/usecase/board"
	userInteractor "todone/pkg/api/usecase/user"
	"todone/pkg/domain/repository"
	boardSvc "todone/pkg/domain/service/board"
	userSvc "todone/pkg/domain/service/user"
	boardRepo "todone/pkg/infrastructure/mysql/board"
	userRepo "todone/pkg/infrastructure/mysql/user"

	"github.com/google/wire"
)

func InitUserAPI(masterTxManager repository.MasterTxManager) userHandler.Server {
	wire.Build(userRepo.New, userSvc.New, userInteractor.New, userHandler.New)

	return userHandler.Server{}
}

func InitBoardAPI(masterTxManager repository.MasterTxManager) boardHandler.Server {
	wire.Build(boardRepo.New, userRepo.New, userSvc.New, boardSvc.New, boardInteractor.New, boardHandler.New)

	return boardHandler.Server{}
}
