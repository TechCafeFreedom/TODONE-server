package entity

import (
	"todone/db/mysql/model"
)

type Label struct {
	ID    int
	Name  string
	Color int
}

type LabelSlice []*Label

func ConvertToLabelSliceEntity(labelSlice model.LabelSlice) LabelSlice {
	res := make(LabelSlice, 0, len(labelSlice))
	for _, labelData := range labelSlice {
		res = append(res, &Label{
			ID:    labelData.ID,
			Name:  labelData.Name,
			Color: labelData.Color,
		})
	}
	return res
}
