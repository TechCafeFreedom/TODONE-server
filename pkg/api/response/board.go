package response

import (
	"todone/pkg/domain/entity"
)

type MemberResponse struct {
	Member *UserResponse `json:"member"`
	Role   int           `json:"role"`
}

type MembersResponse []*MemberResponse

type BoardResponse struct {
	ID          int             `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Author      *UserResponse   `json:"author"`
	Members     MembersResponse `json:"members"`
	Kanbans     KanbansResponse `json:"kanbans"`
	Labels      LabelsResponse  `json:"available_labels"`
}

type BoardsResponse []*BoardResponse

func ConvertToMembersResponse(memberSlice entity.MemberSlice) MembersResponse {
	// nilチェック
	if len(memberSlice) == 0 {
		return nil
	}

	res := make(MembersResponse, 0, len(memberSlice))
	for _, member := range memberSlice {
		res = append(res, &MemberResponse{
			Member: ConvertToUserResponse(member.Member),
			Role:   member.Role,
		})
	}
	return res
}

func ConvertToBoardResponse(boardData *entity.Board) *BoardResponse {
	return &BoardResponse{
		ID:          boardData.ID,
		Title:       boardData.Title,
		Description: boardData.Description,
		Author:      ConvertToUserResponse(boardData.Author),
		Members:     ConvertToMembersResponse(boardData.Members),
		Kanbans:     ConvertToKanbansResponse(boardData.Kanbans),
		Labels:      ConvertToLabelsResponse(boardData.Labels),
	}
}

func ConvertToBoardsResponse(boardSlice entity.BoardSlice) BoardsResponse {
	res := make(BoardsResponse, 0, len(boardSlice))
	for _, boardData := range boardSlice {
		res = append(res, ConvertToBoardResponse(boardData))
	}
	return res
}
