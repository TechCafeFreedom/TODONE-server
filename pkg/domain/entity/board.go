package entity

type Board struct {
	ID          int
	Title       string
	Description string
	Author      *User
	Members     MemberSlice
	Kanbans     KanbanSlice
	Labels      LabelSlice
}

type BoardSlice []*Board
