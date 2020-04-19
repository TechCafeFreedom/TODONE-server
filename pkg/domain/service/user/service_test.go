package user

import (
	"testing"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/user/mock_user"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	userID    = 1
	uid       = "uid"
	name      = "name"
	thumbnail = "thumbnail"
)

func TestService_CreateNewUser(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().InsertUser(ctx, masterTx, uid, name, thumbnail).Return(nil).Times(1)

	service := New(userRepository)
	err := service.CreateNewUser(ctx, masterTx, uid, name, thumbnail)

	assert.NoError(t, err)
}

func TestService_GetByPK(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()

	existedUser := &entity.User{
		ID:        userID,
		Name:      name,
		Thumbnail: thumbnail,
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

	existedUsers := entity.UserSlice{
		{
			ID:        userID,
			Name:      name,
			Thumbnail: thumbnail,
		},
	}

	userRepository := mock_user.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectAll(ctx, masterTx).Return(existedUsers, nil).Times(1)

	service := New(userRepository)
	users, err := service.GetAll(ctx, masterTx)

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
