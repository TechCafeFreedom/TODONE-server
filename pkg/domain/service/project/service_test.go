package project

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/project/mock_project"

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

func TestService_CreateNewUser(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newProject := &model.Project{
		UserID:      userID,
		Title:       title,
		Description: description,
	}

	masterTx := repository.NewMockMasterTx()

	projectRepository := mock_project.NewMockRepository(ctrl)
	projectRepository.EXPECT().InsertProject(ctx, masterTx, newProject).Return(nil).Times(1)

	service := New(projectRepository)
	err := service.CreateNewProject(ctx, masterTx, userID, title, description)

	assert.NoError(t, err)
}

func TestService_GetByPK(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedProject := &model.Project{
		ID:          id,
		UserID:      userID,
		Title:       title,
		Description: description,
	}

	masterTx := repository.NewMockMasterTx()

	projectRepository := mock_project.NewMockRepository(ctrl)
	projectRepository.EXPECT().SelectByPK(ctx, masterTx, id).Return(existedProject, nil).Times(1)

	service := New(projectRepository)
	projects, err := service.GetByPK(ctx, masterTx, id)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}

func TestService_GetByUserID(t *testing.T) {
	ctx := &gin.Context{}
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

	masterTx := repository.NewMockMasterTx()

	projectRepository := mock_project.NewMockRepository(ctrl)
	projectRepository.EXPECT().SelectByUserID(ctx, masterTx, userID).Return(userProjects, nil).Times(1)

	service := New(projectRepository)
	projects, err := service.GetByUserID(ctx, masterTx, userID)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}

func TestService_SelectAll(t *testing.T) {
	ctx := &gin.Context{}
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

	masterTx := repository.NewMockMasterTx()

	projectRepository := mock_project.NewMockRepository(ctrl)
	projectRepository.EXPECT().SelectAll(ctx, masterTx).Return(existedProjects, nil).Times(1)

	service := New(projectRepository)
	projects, err := service.GetAll(ctx, masterTx)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}
