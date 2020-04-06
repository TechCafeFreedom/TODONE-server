package project

import (
	"testing"
	"todone/db/mysql/model"
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

	projectRepository := mock_project.NewMockRepository(ctrl)
	projectRepository.EXPECT().InsertProject(newProject).Return(nil).Times(1)

	service := New(projectRepository)
	err := service.CreateNewProject(userID, title, description)

	assert.NoError(t, err)
}

func TestService_GetByPK(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedProject := &model.Project{
		ID:          id,
		UserID:      userID,
		Title:       title,
		Description: description,
	}

	projectRepository := mock_project.NewMockRepository(ctrl)
	projectRepository.EXPECT().SelectByPK(id).Return(existedProject, nil).Times(1)

	service := New(projectRepository)
	projects, err := service.GetByPK(id)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}

func TestService_GetByUserID(t *testing.T) {
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

	projectRepository := mock_project.NewMockRepository(ctrl)
	projectRepository.EXPECT().SelectByUserID(userID).Return(userProjects, nil).Times(1)

	service := New(projectRepository)
	projects, err := service.GetByUserID(userID)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
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

	projectRepository := mock_project.NewMockRepository(ctrl)
	projectRepository.EXPECT().SelectAll().Return(existedProjects, nil).Times(1)

	service := New(projectRepository)
	projects, err := service.GetAll()

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}
