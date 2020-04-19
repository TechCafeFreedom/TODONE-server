package entity

import (
	"todone/db/mysql/model"
)

type Kanban struct {
	ID        int
	Title     string
	Author    *User
	Cards     CardSlice
	IsArchive bool
}

type KanbanSlice []*Kanban

func ConvertToKanbanEntity(kanbanData *model.Kanban) *Kanban {
	// 関連テーブル情報が何もない場合はnilを詰めて返却する
	if kanbanData.R == nil {
		return &Kanban{
			ID:        kanbanData.ID,
			Title:     kanbanData.Title,
			IsArchive: kanbanData.IsArchive,
		}
	}

	return &Kanban{
		ID:        kanbanData.ID,
		Title:     kanbanData.Title,
		Author:    ConvertToUserEntity(kanbanData.R.User),
		Cards:     ConvertToCardSliceEntity(kanbanData.R.Cards),
		IsArchive: kanbanData.IsArchive,
	}
}

func ConvertToKanbanSliceEntity(kanbanSlice model.KanbanSlice) KanbanSlice {
	res := make(KanbanSlice, 0, len(kanbanSlice))
	for _, kanbanData := range kanbanSlice {
		res = append(res, ConvertToKanbanEntity(kanbanData))
	}
	return res
}
