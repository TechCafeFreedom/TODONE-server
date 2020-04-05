package project

import (
	"testing"
	model "todone/db/mysql/models"
	"todone/pkg/domain/repository/project/mock_project"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	id          = 1
	userID      = "userID"
	title       = "title"
	description = "description"
)

func TestService_CreateNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newProject := &model.Project{
		UserID:      userID,
		Title:       title,
		Description: description,
	}

	userRepository := mock_project.NewMockRepository(ctrl)
	userRepository.EXPECT().InsertProject(newProject).Return(nil).Times(1)

	service := New(userRepository)
	err := service.CreateNewProject(userID, title, description)

	assert.NoError(t, err)
}

func TestService_SelectAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedProjects := model.ProjectSlice{
		{
			ID:          id,
			UserID:      userID,
			Title:       title,
			Description: description,
		},
	}

	userRepository := mock_project.NewMockRepository(ctrl)
	userRepository.EXPECT().SelectAll().Return(existedProjects, nil).Times(1)

	service := New(userRepository)
	users, err := service.SelectAll()

	assert.NoError(t, err)
	assert.NotNil(t, users)
}
