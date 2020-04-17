package board

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
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
func (p *boardRepositoryImpliment) InsertBoard(ctx *gin.Context, masterTx repository.MasterTx, boardData *model.Board) error {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return err
	}
	if err := boardData.Insert(ctx, exec, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// プロジェクト取得機能（PK: id）
func (p *boardRepositoryImpliment) SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*model.Board, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	boardData, err := model.Boards(
		qm.Load(model.BoardRels.UsersBoards),
		qm.Load(model.UsersBoardRels.User),
		qm.Load(model.BoardRels.Kanbans),
		qm.Load(model.BoardRels.Labels),
		model.BoardWhere.ID.EQ(id),
	).One(ctx, exec)
	if err != nil {
		return nil, err
	}

	return boardData, nil
}

// ユーザのもつプロジェクト取得機能
func (p *boardRepositoryImpliment) SelectByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID int) (model.BoardSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	boards, err := model.Boards(
		qm.Load(model.BoardRels.User),
		model.BoardWhere.UserID.EQ(userID),
	).All(ctx, exec)
	if err != nil {
		return nil, err
	}

	return boards, nil
}

// プロジェクト全件取得機能
func (p *boardRepositoryImpliment) SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (model.BoardSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	queries := []qm.QueryMod{}
	boards, err := model.Boards(queries...).All(ctx, exec)
	if err != nil {
		return nil, err
	}

	return boards, nil
}
