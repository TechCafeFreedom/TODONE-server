package board

import (
	"context"
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board"
	"todone/pkg/infrastructure/mysql"
	"todone/pkg/terrors"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type boardRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) board.Repository {
	return &boardRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

// プロジェクト作成機能
func (p *boardRepositoryImpliment) InsertBoard(ctx context.Context, masterTx repository.MasterTx, userID int, title, description string) error {
	newBoardData := &model.Board{
		UserID:      userID,
		Title:       title,
		Description: null.StringFrom(description),
	}

	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return terrors.Stack(err)
	}
	if err := newBoardData.Insert(ctx, exec, boil.Infer()); err != nil {
		return terrors.Stack(err)
	}

	return nil
}

// プロジェクト取得機能（PK: id）
func (p *boardRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*entity.Board, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	boardData, err := model.Boards(
		qm.Load(model.BoardRels.UsersBoards),
		qm.Load(model.UsersBoardRels.User),
		qm.Load(model.BoardRels.Kanbans),
		qm.Load(model.BoardRels.Labels),
		model.BoardWhere.ID.EQ(id),
	).One(ctx, exec)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	return entity.ConvertToBoardEntity(boardData), nil
}

// ユーザのもつプロジェクト取得機能
func (p *boardRepositoryImpliment) SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	boards, err := model.Boards(
		qm.Load(model.BoardRels.User),
		model.BoardWhere.UserID.EQ(userID),
	).All(ctx, exec)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	return entity.ConvertToBoardSliceEntity(boards), nil
}

// プロジェクト全件取得機能
func (p *boardRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.BoardSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	queries := []qm.QueryMod{}
	boards, err := model.Boards(queries...).All(ctx, exec)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	return entity.ConvertToBoardSliceEntity(boards), nil
}
