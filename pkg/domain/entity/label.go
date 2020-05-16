package entity

type Label struct {
	ID    int
	Name  string
	Color int
}

type LabelSlice []*Label
