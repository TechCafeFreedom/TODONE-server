package project

import (
	"testing"
	model "todone/db/mysql/models"
	"todone/pkg/domain/service/project/mock_project"

	"github.com/gin-gonic/gin"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	id          = 1
	userID      = "userID"
	title       = "title"
	description = "description"
)

func TestIntereractor_CreateNewUser(t *testing.T) {
	ctx := &gin.Context{}
	ctx.Set("AUTHED_USER_ID", userID)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userService := mock_project.NewMockService(ctrl)
	userService.EXPECT().CreateNewProject(userID, title, description).Return(nil).Times(1)

	interactor := New(userService)
	err := interactor.CreateNewProject(ctx, title, description)

	assert.NoError(t, err)
}

func TestIntereractor_SelectAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedUsers := model.ProjectSlice{
		{
			ID:          id,
			UserID:      userID,
			Title:       title,
			Description: description,
		},
	}

	userService := mock_project.NewMockService(ctrl)
	userService.EXPECT().SelectAll().Return(existedUsers, nil).Times(1)

	interactor := New(userService)
	users, err := interactor.SelectAll()

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
