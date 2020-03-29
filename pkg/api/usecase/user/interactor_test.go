package user

import (
	"fmt"
	"testing"
	model "todone/db/mysql/models"
	"todone/pkg/domain/service/user/mock_user"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// uidなどの実行ごとにランダムな値に変わってしまう値が含まれる場合は、gomockのEXPECT関数を拡張する必要がある。
type idMatcher struct {
	user *model.User
	id   string
}

func (e idMatcher) Matches(x interface{}) bool {
	cast, ok := x.(*model.User)
	if !ok {
		return false
	}
	if cast.ID == "" {
		return false
	}
	return true
}

func (e idMatcher) String() string {
	return fmt.Sprintf("is equal to %v", e.id)
}

const (
	id        = "id"
	userID    = "userID"
	email     = "email"
	password  = "password"
	name      = "name"
	thumbnail = "thumbnail"
)

func TestIntereractor_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	loginedUser := &model.User{
		ID:        id,
		UserID:    userID,
		Email:     email,
		Password:  password,
		Name:      name,
		Thumbnail: thumbnail,
	}

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().Login(email, password).Return(loginedUser, nil).Times(1)

	interactor := New(userService)
	authedUser, err := interactor.Login(email, password)

	assert.NoError(t, err)
	assert.NotNil(t, authedUser)
}

func TestIntereractor_GetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUsers := model.UserSlice{
		{
			ID:        id,
			UserID:    userID,
			Email:     email,
			Password:  password,
			Name:      name,
			Thumbnail: thumbnail,
		},
	}

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetUser(id).Return(existedUsers, nil).Times(1)

	interactor := New(userService)
	users, err := interactor.GetUser(id)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestIntereractor_CreateNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newUser := &model.User{
		ID:        id,
		UserID:    userID,
		Email:     email,
		Password:  password,
		Name:      name,
		Thumbnail: thumbnail,
	}

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().CreateNewUser(idMatcher{
		user: newUser,
		id:   id,
	}).Return(nil).Times(1)

	interactor := New(userService)
	token, err := interactor.CreateNewUser(userID, name, thumbnail, email, password)

	assert.NoError(t, err)
	assert.NotNil(t, token)
}

func TestIntereractor_SelectAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUsers := model.UserSlice{
		{
			ID:        id,
			UserID:    userID,
			Email:     email,
			Password:  password,
			Name:      name,
			Thumbnail: thumbnail,
		},
	}

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().SelectAll().Return(existedUsers, nil).Times(1)

	interactor := New(userService)
	users, err := interactor.SelectAll()

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
