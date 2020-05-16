package label

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
)

func ConvertToLabelEntity(labelData *model.Label) *entity.Label {
	return &entity.Label{
		ID:    labelData.ID,
		Name:  labelData.Name,
		Color: labelData.Color,
	}
}

func ConvertToLabelSliceEntity(labelSlice model.LabelSlice) entity.LabelSlice {
	res := make(entity.LabelSlice, 0, len(labelSlice))
	for _, labelData := range labelSlice {
		res = append(res, ConvertToLabelEntity(labelData))
	}
	return res
}
