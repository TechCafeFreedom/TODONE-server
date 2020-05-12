package board

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board"
	"todone/pkg/infrastructure/mysql"
	"todone/pkg/infrastructure/mysql/kanban"
	"todone/pkg/infrastructure/mysql/label"
	"todone/pkg/infrastructure/mysql/user"
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
func (b *boardRepositoryImpliment) InsertBoard(ctx context.Context, masterTx repository.MasterTx, userID int, title, description string) error {
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
func (b *boardRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*entity.Board, error) {
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
		log.Println("Error occred when DB access.")
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return ConvertToBoardEntity(boardData), nil
}

// ユーザのもつプロジェクト取得機能
func (b *boardRepositoryImpliment) SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error) {
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
		log.Println("Error occred when DB access.")
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return ConvertToBoardSliceEntity(boards), nil
}

// プロジェクト全件取得機能
func (b *boardRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.BoardSlice, error) {
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
		log.Println("Error occred when DB access.")
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "サーバでエラーが発生しました。", "Error occured at server.")
	}

	return ConvertToBoardSliceEntity(boards), nil
}

func ConvertToBoardEntity(boardData *model.Board) *entity.Board {
	// 関連テーブル情報が何もない場合はnilを詰めて返却する
	if boardData.R == nil {
		return &entity.Board{
			ID:          boardData.ID,
			Title:       boardData.Title,
			Description: boardData.Description.String,
		}
	}

	return &entity.Board{
		ID:          boardData.ID,
		Title:       boardData.Title,
		Description: boardData.Description.String,
		Author:      user.ConvertToUserEntity(boardData.R.User),
		Members:     ConvertToMemberSliceEntity(boardData.R.UsersBoards),
		Kanbans:     kanban.ConvertToKanbanSliceEntity(boardData.R.Kanbans),
		Labels:      label.ConvertToLabelSliceEntity(boardData.R.Labels),
	}
}

func ConvertToBoardSliceEntity(boardSlice model.BoardSlice) entity.BoardSlice {
	res := make(entity.BoardSlice, 0, len(boardSlice))
	for _, boardData := range boardSlice {
		res = append(res, ConvertToBoardEntity(boardData))
	}
	return res
}

func ConvertToMemberEntity(userBoardData *model.UsersBoard) *entity.Member {
	return &entity.Member{
		Member: user.ConvertToUserEntity(userBoardData.R.User),
		Role:   int(userBoardData.Role),
	}
}

func ConvertToMemberSliceEntity(userBoardSlice model.UsersBoardSlice) entity.MemberSlice {
	res := make(entity.MemberSlice, 0, len(userBoardSlice))
	for _, userBoard := range userBoardSlice {
		res = append(res, ConvertToMemberEntity(userBoard))
	}
	return res
}
