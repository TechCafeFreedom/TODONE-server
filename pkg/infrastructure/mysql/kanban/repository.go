package kanban

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/infrastructure/mysql/card"
	"todone/pkg/infrastructure/mysql/user"
)

func ConvertToKanbanEntity(kanbanData *model.Kanban) *entity.Kanban {
	// 関連テーブル情報が何もない場合はnilを詰めて返却する
	if kanbanData.R == nil {
		return &entity.Kanban{
			ID:        kanbanData.ID,
			Title:     kanbanData.Title,
			IsArchive: kanbanData.IsArchive,
		}
	}

	return &entity.Kanban{
		ID:        kanbanData.ID,
		Title:     kanbanData.Title,
		Author:    user.ConvertToUserEntity(kanbanData.R.User),
		Cards:     card.ConvertToCardSliceEntity(kanbanData.R.Cards),
		IsArchive: kanbanData.IsArchive,
	}
}

func ConvertToKanbanSliceEntity(kanbanSlice model.KanbanSlice) entity.KanbanSlice {
	res := make(entity.KanbanSlice, 0, len(kanbanSlice))
	for _, kanbanData := range kanbanSlice {
		res = append(res, ConvertToKanbanEntity(kanbanData))
	}
	return res
}
