package user

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/user/mock_user"

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

func TestService_CreateNewUser(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	newUser := &model.User{
		AccessToken: accessToken,
		Name:        name,
		Thumbnail:   null.StringFrom(thumbnail),
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().InsertUser(ctx, masterTx, newUser).Return(nil).Times(1)

	service := New(userRepository)
	err := service.CreateNewUser(ctx, masterTx, accessToken, name, thumbnail)

	assert.NoError(t, err)
}

func TestService_GetByPK(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	existedUser := &model.User{
		ID:          userID,
		AccessToken: accessToken,
		Name:        name,
		Thumbnail:   null.StringFrom(thumbnail),
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

	masterTx := repository.NewMockMasterTx()

	existedUsers := model.UserSlice{
		{
			ID:          userID,
			AccessToken: accessToken,
			Name:        name,
			Thumbnail:   null.StringFrom(thumbnail),
		},
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectAll(ctx, masterTx).Return(existedUsers, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetAll(ctx, masterTx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
