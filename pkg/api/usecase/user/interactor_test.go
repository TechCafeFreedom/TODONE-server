package user

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/service/user/mock_user"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null"
)

const (
	userID      = 1
	accessToken = "accessToken"
	name        = "name"
	thumbnail   = "thumbnail"
)

func TestIntereractor_CreateNewUser(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().CreateNewUser(ctx, masterTx, accessToken, name, thumbnail).Return(nil).Times(1)

	interactor := New(masterTxManager, userService)
	err := interactor.CreateNewUser(ctx, accessToken, name, thumbnail)

	assert.NoError(t, err)
}

func TestIntereractor_GetUserProfile(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUser := &model.User{
		ID:          userID,
		AccessToken: accessToken,
		Name:        name,
		Thumbnail:   null.StringFrom(thumbnail),
	}

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetByAccessToken(ctx, masterTx, accessToken).Return(existedUser, nil).Times(1)

	interactor := New(masterTxManager, userService)
	users, err := interactor.GetUserProfile(ctx, accessToken)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}

func TestIntereractor_GetAll(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUsers := model.UserSlice{
		{
			ID:          userID,
			AccessToken: accessToken,
			Name:        name,
			Thumbnail:   null.StringFrom(thumbnail),
		},
	}

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetAll(ctx, masterTx).Return(existedUsers, nil).Times(1)

	interactor := New(masterTxManager, userService)
	users, err := interactor.GetAll(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
