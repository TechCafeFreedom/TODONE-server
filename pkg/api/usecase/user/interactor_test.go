package user

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/service/user/mock_user"
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

func TestIntereractor_CreateNewUser(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := mysql.NewMockMasterTx()
	masterTxManager := mysql.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().CreateNewUser(userID, name, thumbnail).Return(nil).Times(1)

	interactor := New(masterTxManager, userService)
	err := interactor.CreateNewUser(ctx, userID, name, thumbnail)

	assert.NoError(t, err)
}

func TestIntereractor_GetByPK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUser := &model.User{
		UserID:    userID,
		Name:      name,
		Thumbnail: null.StringFrom(thumbnail),
	}

	masterTx := mysql.NewMockMasterTx()
	masterTxManager := mysql.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetByPK(userID).Return(existedUser, nil).Times(1)

	interactor := New(masterTxManager, userService)
	users, err := interactor.GetByPK(userID)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestIntereractor_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUsers := model.UserSlice{
		{
			UserID:    userID,
			Name:      name,
			Thumbnail: null.StringFrom(thumbnail),
		},
	}

	masterTx := mysql.NewMockMasterTx()
	masterTxManager := mysql.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetAll().Return(existedUsers, nil).Times(1)

	interactor := New(masterTxManager, userService)
	users, err := interactor.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
