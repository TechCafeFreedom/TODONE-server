package response

import (
	"todone/pkg/domain/entity"
)

type LabelResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color int    `json:"color"`
}

type LabelsResponse []*LabelResponse

func ConvertToLabelsResponse(labelSlice entity.LabelSlice) LabelsResponse {
	// nilチェック
	if len(labelSlice) == 0 {
		return nil
	}

	res := make(LabelsResponse, 0, len(labelSlice))
	for _, labelData := range labelSlice {
		res = append(res, &LabelResponse{
			ID:    labelData.ID,
			Name:  labelData.Name,
			Color: labelData.Color,
		})
	}
	return res
}
