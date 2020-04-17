package entity

import (
	"todone/db/mysql/model"
)

type Member struct {
	Member *User
	Role   int
}

type MemberSlice []*Member

func ConvertToMemberSliceEntity(userBoardSlice model.UsersBoardSlice) MemberSlice {
	res := make(MemberSlice, 0, len(userBoardSlice))
	for _, userBoard := range userBoardSlice {
		res = append(res, &Member{
			Member: ConvertToUserEntity(userBoard.R.User),
			Role:   int(userBoard.Role),
		})
	}
	return res
}
