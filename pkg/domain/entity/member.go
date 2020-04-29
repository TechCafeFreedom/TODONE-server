package entity

type Member struct {
	Member *User
	Role   int
}

type MemberSlice []*Member
