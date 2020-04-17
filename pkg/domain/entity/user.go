package entity

import (
	"todone/db/mysql/model"
)

type User struct {
	ID        int
	Name      string
	Thumbnail string
}

type UserSlice []*User

func ConvertToUserEntity(userData *model.User) *User {
	return &User{
		Name:      userData.Name,
		Thumbnail: userData.Thumbnail.String,
	}
}
