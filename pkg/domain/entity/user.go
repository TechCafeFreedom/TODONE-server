package entity

type User struct {
	ID        int
	Name      string
	Thumbnail string
}

type UserSlice []*User
