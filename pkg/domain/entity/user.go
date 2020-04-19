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

func ConvertToUserSliceEntity(userSlice model.UserSlice) UserSlice {
	res := make(UserSlice, 0, len(userSlice))
	for _, userData := range userSlice {
		res = append(res, ConvertToUserEntity(userData))
	}
	return res
}

func ConvertToUserEntity(userData *model.User) *User {
	return &User{
		ID:        userData.ID,
		Name:      userData.Name,
		Thumbnail: userData.Thumbnail.String,
	}
}
