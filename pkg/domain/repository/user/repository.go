package user

import "todone/db/mysql/model"

type Repository interface {
	InsertUser(user *model.User) error
	SelectByPK(userID string) (*model.User, error)
	SelectAll() (model.UserSlice, error)
}
