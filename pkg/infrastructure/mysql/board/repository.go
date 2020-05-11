package board

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board"
	"todone/pkg/infrastructure/mysql"
	"todone/pkg/terrors"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	if err == sql.ErrNoRows {
		messageJP := fmt.Sprintf("指定されたボードは見つかりませんでした。board_id: %v", id)
		messageEN := fmt.Sprintf("This ID's board doesn't exists.board_id: %v", id)
		return nil, terrors.Newf(http.StatusInternalServerError, messageJP, messageEN)
	} else if err != nil {
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "DBアクセス時にエラーが発生しました。", "Error occured in DB access.")
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
	if err == sql.ErrNoRows {
		messageJP := fmt.Sprintf("ユーザの所有するボードは１件もありませんでした。")
		messageEN := fmt.Sprintf("User's board doesn't exists.")
		return nil, terrors.Newf(http.StatusInternalServerError, messageJP, messageEN)
	} else if err != nil {
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "DBアクセス時にエラーが発生しました。", "Error occured in DB access.")
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
	if err == sql.ErrNoRows {
		messageJP := fmt.Sprintf("ボードは１件もありませんでした。")
		messageEN := fmt.Sprintf("Board doesn't exists.")
		return nil, terrors.Newf(http.StatusInternalServerError, messageJP, messageEN)
	} else if err != nil {
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "DBアクセス時にエラーが発生しました。", "Error occured in DB access.")
	}

	return entity.ConvertToBoardSliceEntity(boards), nil
}
