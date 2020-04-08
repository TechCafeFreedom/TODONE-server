package user

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/repository/user/mock_user"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
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
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := mysql.NewMockMasterTx()

	newUser := &model.User{
		UserID:    userID,
		Name:      name,
		Thumbnail: null.StringFrom(thumbnail),
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().InsertUser(ctx, masterTx, newUser).Return(nil).Times(1)

	service := New(userRepository)
	err := service.CreateNewUser(ctx, masterTx, userID, name, thumbnail)

	assert.NoError(t, err)
}

func TestService_GetByPK(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := mysql.NewMockMasterTx()

	existedUser := &model.User{
		UserID:    userID,
		Name:      name,
		Thumbnail: null.StringFrom(thumbnail),
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectByPK(ctx, masterTx, userID).Return(existedUser, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetByPK(ctx, masterTx, userID)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestService_SelectAll(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := mysql.NewMockMasterTx()

	existedUsers := model.UserSlice{
		{
			UserID:    userID,
			Name:      name,
			Thumbnail: null.StringFrom(thumbnail),
		},
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectAll(ctx, masterTx).Return(existedUsers, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetAll(ctx, masterTx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
