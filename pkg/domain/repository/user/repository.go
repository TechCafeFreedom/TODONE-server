package user

import (
	model "todone/db/mysql/models"
)

type Repository interface {
	InsertUser(user *model.User) error
	Login(email, password string) (*model.User, error)
	UpdateImg(s3Url string) error
	SelectByUserId(id string) ([]*model.User, error)
	SelectAll() ([]*model.User, error)
	SelectByEmail(email string) ([]*model.User, error)
}
