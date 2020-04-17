package response

import (
	"todone/pkg/domain/entity"
)

type KanbanResponse struct {
	ID        int           `json:"id"`
	Title     string        `json:"title"`
	Author    *UserResponse `json:"author"`
	Cards     CardsResponse `json:"cards"`
	IsArchive bool          `json:"is_archive"`
}

type KanbansResponse []*KanbanResponse

func ConvertToKanbansResponse(kanbanSlice entity.KanbanSlice) KanbansResponse {
	res := make(KanbansResponse, 0, len(kanbanSlice))
	for _, kanbanData := range kanbanSlice {
		res = append(res, &KanbanResponse{
			ID:        kanbanData.ID,
			Title:     kanbanData.Title,
			Author:    ConvertToUserResponse(kanbanData.Author),
			Cards:     ConvertToCardsResponse(kanbanData.Cards),
			IsArchive: false,
		})
	}
	return res
}
