package entity

import (
	"todone/db/mysql/model"
)

type Board struct {
	ID          int
	Title       string
	Description string
	Author      *User
	Members     MemberSlice
	Kanbans     KanbanSlice
	Labels      LabelSlice
}

type BoardSlice []*Board

func ConvertToBoardEntity(boardData *model.Board) *Board {
	// 関連テーブル情報が何もない場合はnilを詰めて返却する
	if boardData.R == nil {
		return &Board{
			ID:          boardData.ID,
			Title:       boardData.Title,
			Description: boardData.Description.String,
		}
	}

	return &Board{
		ID:          boardData.ID,
		Title:       boardData.Title,
		Description: boardData.Description.String,
		Author:      ConvertToUserEntity(boardData.R.User),
		Members:     ConvertToMemberSliceEntity(boardData.R.UsersBoards),
		Kanbans:     ConvertToKanbanSliceEntity(boardData.R.Kanbans),
		Labels:      ConvertToLabelSliceEntity(boardData.R.Labels),
	}
}

func ConvertToBoardSliceEntity(boardSlice model.BoardSlice) BoardSlice {
	res := make(BoardSlice, 0, len(boardSlice))
	for _, boardData := range boardSlice {
		res = append(res, ConvertToBoardEntity(boardData))
	}
	return res
}
