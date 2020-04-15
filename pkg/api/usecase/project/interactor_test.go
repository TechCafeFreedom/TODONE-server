package project

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
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

func TestIntereractor_CreateNewProject(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	projectService := mock_project.NewMockService(ctrl)
	projectService.EXPECT().CreateNewProject(ctx, masterTx, userID, title, description).Return(nil).Times(1)

	interactor := New(masterTxManager, projectService)
	err := interactor.CreateNewProject(ctx, userID, title, description)

	assert.NoError(t, err)
}

func TestIntereractor_GetByPK(t *testing.T) {
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
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	projectService := mock_project.NewMockService(ctrl)
	projectService.EXPECT().GetByPK(ctx, masterTx, id).Return(existedProject, nil).Times(1)

	interactor := New(masterTxManager, projectService)
	projects, err := interactor.GetByPK(ctx, id)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}

func TestIntereractor_GetByUserID(t *testing.T) {
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
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	projectService := mock_project.NewMockService(ctrl)
	projectService.EXPECT().GetByUserID(ctx, masterTx, userID).Return(userProjects, nil).Times(1)

	interactor := New(masterTxManager, projectService)
	projects, err := interactor.GetByUserID(ctx, userID)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}

func TestIntereractor_GetAll(t *testing.T) {
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
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	projectService := mock_project.NewMockService(ctrl)
	projectService.EXPECT().GetAll(ctx, masterTx).Return(existedProjects, nil).Times(1)

	interactor := New(masterTxManager, projectService)
	projects, err := interactor.GetAll(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, projects)
}
