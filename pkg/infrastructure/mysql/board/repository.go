package board

import (
	"context"
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board"
	"todone/pkg/infrastructure/mysql"
	"todone/pkg/infrastructure/mysql/kanban"
	"todone/pkg/infrastructure/mysql/label"
	"todone/pkg/infrastructure/mysql/user"

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
func (b *boardRepositoryImpliment) InsertBoard(ctx context.Context, masterTx repository.MasterTx, userID int, title, description string) error {
	newBoardData := &model.Board{
		UserID:      userID,
		Title:       title,
		Description: null.StringFrom(description),
	}

	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return err
	}
	if err := newBoardData.Insert(ctx, exec, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// プロジェクト取得機能（PK: id）
func (b *boardRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*entity.Board, error) {
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

	return ConvertToBoardEntity(boardData), nil
}

// ユーザのもつプロジェクト取得機能
func (b *boardRepositoryImpliment) SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error) {
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

	return ConvertToBoardSliceEntity(boards), nil
}

// プロジェクト全件取得機能
func (b *boardRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.BoardSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	queries := []qm.QueryMod{}
	boards, err := model.Boards(queries...).All(ctx, exec)
	if err != nil {
		return nil, err
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
