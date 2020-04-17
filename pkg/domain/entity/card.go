package entity

import (
	"todone/db/mysql/model"
)

type Card struct {
	ID        int
	Title     string
	Content   string
	Deadline  int
	Author    *User
	Labels    LabelSlice
	IsArchive bool
}

type CardSlice []*Card

func ConvertToCardEntity(cardData *model.Card) *Card {
	// 関連テーブル情報が何もない場合はnilを詰めて返却する
	if cardData.R == nil {
		return &Card{
			ID:        cardData.ID,
			Title:     cardData.Title,
			Content:   cardData.Content.String,
			Deadline:  int(cardData.Deadline.Time.Unix()),
			IsArchive: cardData.IsArchive,
		}
	}

	return &Card{
		ID:        cardData.ID,
		Title:     cardData.Title,
		Content:   cardData.Content.String,
		Deadline:  int(cardData.Deadline.Time.Unix()),
		Author:    ConvertToUserEntity(cardData.R.User),
		IsArchive: cardData.IsArchive,
		Labels:    ConvertToLabelSliceEntity(cardData.R.Labels),
	}
}

func ConvertToCardSliceEntity(cardSlice model.CardSlice) CardSlice {
	res := make(CardSlice, 0, len(cardSlice))
	for _, cardData := range cardSlice {
		res = append(res, ConvertToCardEntity(cardData))
	}
	return res
}
