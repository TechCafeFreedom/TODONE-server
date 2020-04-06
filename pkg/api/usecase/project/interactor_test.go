package project

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/service/project/mock_project"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	id          = 1
	userID      = "userID"
	title       = "title"
	description = "description"
)

func TestIntereractor_CreateNewProject(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectService := mock_project.NewMockService(ctrl)
	projectService.EXPECT().CreateNewProject(userID, title, description).Return(nil).Times(1)

	interactor := New(projectService)
	err := interactor.CreateNewProject(userID, title, description)

	assert.NoError(t, err)
}

func TestIntereractor_GetByPK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedProject := &model.Project{
		ID:          id,
		UserID:      userID,
		Title:       title,
		Description: description,
	}

	projectService := mock_project.NewMockService(ctrl)
	projectService.EXPECT().GetByPK(id).Return(existedProject, nil).Times(1)

	interactor := New(projectService)
	projects, err := interactor.GetByPK(id)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}

func TestIntereractor_GetByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userProjects := model.ProjectSlice{
		{
			ID:          id,
			UserID:      userID,
			Title:       title,
			Description: description,
		},
	}

	projectService := mock_project.NewMockService(ctrl)
	projectService.EXPECT().GetByUserID(userID).Return(userProjects, nil).Times(1)

	interactor := New(projectService)
	projects, err := interactor.GetByUserID(userID)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}

func TestIntereractor_GetAll(t *testing.T) {
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

	projectService := mock_project.NewMockService(ctrl)
	projectService.EXPECT().GetAll().Return(existedProjects, nil).Times(1)

	interactor := New(projectService)
	projects, err := interactor.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}
