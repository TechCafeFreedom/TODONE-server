package response

import (
	"todone/pkg/domain/entity"
)

type CardResponse struct {
	ID        int            `json:"id"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Deadline  int            `json:"deadline"`
	Author    *UserResponse  `json:"author"`
	IsArchive bool           `json:"is_archive"`
	Labels    LabelsResponse `json:"labels"`
}

type CardsResponse []*CardResponse

func ConvertToCardsResponse(cardSlice entity.CardSlice) CardsResponse {
	res := make(CardsResponse, 0, len(cardSlice))
	for _, cardData := range cardSlice {
		res = append(res, &CardResponse{
			ID:        cardData.ID,
			Title:     cardData.Title,
			Content:   cardData.Content,
			Deadline:  cardData.Deadline,
			Author:    ConvertToUserResponse(cardData.Author),
			IsArchive: cardData.IsArchive,
			Labels:    ConvertToLabelsResponse(cardData.Labels),
		})
	}
	return res
}
