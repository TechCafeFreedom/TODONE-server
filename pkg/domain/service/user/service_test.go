package user

import (
	"testing"
	model "todone/db/mysql/models"
	"todone/pkg/domain/repository/user/mock_user"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	id        = "id"
	userID    = "userID"
	email     = "email"
	password  = "password"
	name      = "name"
	thumbnail = "thumbnail"
)

func TestService_Login(t *testing.T) {
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

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().Login(email, password).Return(loginedUser, nil).Times(1)

	service := New(userRepository)
	authedUser, err := service.Login(email, password)

	assert.NoError(t, err)
	assert.NotNil(t, authedUser)
}

func TestService_GetUser(t *testing.T) {
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

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectByUserID(userID).Return(existedUsers, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetUser(userID)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestService_CreateNewUser(t *testing.T) {
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

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().InsertUser(newUser).Return(nil).Times(1)

	service := New(userRepository)
	err := service.CreateNewUser(newUser)

	assert.NoError(t, err)
}

func TestService_SelectAll(t *testing.T) {
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

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectAll().Return(existedUsers, nil).Times(1)

	service := New(userRepository)
	users, err := service.SelectAll()

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
