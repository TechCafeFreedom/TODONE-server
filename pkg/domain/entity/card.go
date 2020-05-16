package entity

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
