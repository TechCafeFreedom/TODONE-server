package card

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/infrastructure/mysql/label"
	"todone/pkg/infrastructure/mysql/user"
)

func ConvertToCardEntity(cardData *model.Card) *entity.Card {
	// 関連テーブル情報が何もない場合はnilを詰めて返却する
	if cardData.R == nil {
		return &entity.Card{
			ID:        cardData.ID,
			Title:     cardData.Title,
			Content:   cardData.Content.String,
			Deadline:  int(cardData.Deadline.Time.Unix()),
			IsArchive: cardData.IsArchive,
		}
	}

	return &entity.Card{
		ID:        cardData.ID,
		Title:     cardData.Title,
		Content:   cardData.Content.String,
		Deadline:  int(cardData.Deadline.Time.Unix()),
		Author:    user.ConvertToUserEntity(cardData.R.User),
		IsArchive: cardData.IsArchive,
		Labels:    label.ConvertToLabelSliceEntity(cardData.R.Labels),
	}
}

func ConvertToCardSliceEntity(cardSlice model.CardSlice) entity.CardSlice {
	res := make(entity.CardSlice, 0, len(cardSlice))
	for _, cardData := range cardSlice {
		res = append(res, ConvertToCardEntity(cardData))
	}
	return res
}
