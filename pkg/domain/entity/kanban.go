package entity

type Kanban struct {
	ID        int
	Title     string
	Author    *User
	Cards     CardSlice
	IsArchive bool
}

type KanbanSlice []*Kanban
