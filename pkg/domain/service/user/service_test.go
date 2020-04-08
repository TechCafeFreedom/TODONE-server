package user

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/repository/user/mock_user"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"
)

const (
	userID    = "userID"
	name      = "name"
	thumbnail = "thumbnail"
)

func TestService_CreateNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newUser := &model.User{
		UserID:    userID,
		Name:      name,
		Thumbnail: null.StringFrom(thumbnail),
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().InsertUser(newUser).Return(nil).Times(1)

	service := New(userRepository)
	err := service.CreateNewUser(userID, name, thumbnail)

	assert.NoError(t, err)
}

func TestService_GetByPK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUser := &model.User{
		UserID:    userID,
		Name:      name,
		Thumbnail: null.StringFrom(thumbnail),
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectByPK(userID).Return(existedUser, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetByPK(userID)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestService_SelectAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUsers := model.UserSlice{
		{
			UserID:    userID,
			Name:      name,
			Thumbnail: null.StringFrom(thumbnail),
		},
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectAll().Return(existedUsers, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
