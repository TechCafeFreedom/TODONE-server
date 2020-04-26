package board

import (
	"context"
	"net/http"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	boardservice "todone/pkg/domain/service/board"
	userservice "todone/pkg/domain/service/user"
	"todone/pkg/terrors"
)

type Interactor interface {
	CreateNewBoard(ctx context.Context, uid string, title, description string) error
	GetBoardDetail(ctx context.Context, id int) (*entity.Board, error)
	GetUserBoards(ctx context.Context, uid string) (entity.BoardSlice, error)
	GetAll(ctx context.Context) (entity.BoardSlice, error)
}

type intereractor struct {
	masterTxManager repository.MasterTxManager
	boardService    boardservice.Service
	userService     userservice.Service
}

func New(masterTxManager repository.MasterTxManager, boardService boardservice.Service, userService userservice.Service) Interactor {
	return &intereractor{
		masterTxManager: masterTxManager,
		boardService:    boardService,
		userService:     userService,
	}
}

func (i *intereractor) CreateNewBoard(ctx context.Context, uid string, title, description string) error {
	// titleの空文字チェック
	if title == "" {
		return terrors.Newf(http.StatusBadRequest, "ボードタイトルは必須項目です。", "Board title is required.")
	}

	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザのID取得
		userData, err := i.userService.GetByUID(ctx, masterTx, uid)
		if err != nil {
			return terrors.Stack(err)
		}
		// ユーザ存在チェック
		if userData == nil || userData.ID <= 0 {
			return terrors.Newf(http.StatusBadRequest, "ユーザが見つかりませんでした。", "User not found.")
		}

		// 新規ボードの作成
		if err := i.boardService.CreateNewBoard(ctx, masterTx, userData.ID, title, description); err != nil {
			return terrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return terrors.Stack(err)
	}
	return nil
}

func (i *intereractor) GetBoardDetail(ctx context.Context, id int) (*entity.Board, error) {
	var boardData *entity.Board
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ボード詳細情報取得
		boardData, err = i.boardService.GetByPK(ctx, masterTx, id)
		if err != nil {
			return terrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return boardData, nil
}

func (i *intereractor) GetUserBoards(ctx context.Context, uid string) (entity.BoardSlice, error) {
	var boardSlice entity.BoardSlice

	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザのID取得
		userData, err := i.userService.GetByUID(ctx, masterTx, uid)
		if err != nil {
			return terrors.Stack(err)
		}
		// ユーザ所有のボード一覧取得
		boardSlice, err = i.boardService.GetByUserID(ctx, masterTx, userData.ID)
		if err != nil {
			return terrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return boardSlice, nil
}

func (i *intereractor) GetAll(ctx context.Context) (entity.BoardSlice, error) {
	var boardSlice entity.BoardSlice
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// (管理者用)ボード全件取得
		boardSlice, err = i.boardService.GetAll(ctx, masterTx)
		if err != nil {
			return terrors.Stack(err)
		}
		return nil
	})
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return boardSlice, nil
}
