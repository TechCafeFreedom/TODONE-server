package project

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	projectservice "todone/pkg/domain/service/project"

	"github.com/gin-gonic/gin"
)

type Interactor interface {
	CreateNewProject(ctx *gin.Context, userID, title, description string) error
	GetByPK(ctx *gin.Context, id int) (*model.Project, error)
	GetByUserID(ctx *gin.Context, userID string) (model.ProjectSlice, error)
	GetAll(ctx *gin.Context) (model.ProjectSlice, error)
}

type intereractor struct {
	masterTxManager repository.MasterTxManager
	projectService  projectservice.Service
}

func New(masterTxManager repository.MasterTxManager, projectService projectservice.Service) Interactor {
	return &intereractor{
		masterTxManager: masterTxManager,
		projectService:  projectService,
	}
}

func (i *intereractor) CreateNewProject(ctx *gin.Context, userID, title, description string) error {
	err := i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx repository.MasterTx) error {
		if err := i.projectService.CreateNewProject(ctx, masterTx, userID, title, description); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *intereractor) GetByPK(ctx *gin.Context, id int) (*model.Project, error) {
	var projectData *model.Project
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx repository.MasterTx) error {
		projectData, err = i.projectService.GetByPK(ctx, masterTx, id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return projectData, nil
}

func (i *intereractor) GetByUserID(ctx *gin.Context, userID string) (model.ProjectSlice, error) {
	var projectSlice model.ProjectSlice
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx repository.MasterTx) error {
		projectSlice, err = i.projectService.GetByUserID(ctx, masterTx, userID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return projectSlice, nil
}

func (i *intereractor) GetAll(ctx *gin.Context) (model.ProjectSlice, error) {
	var projectSlice model.ProjectSlice
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx repository.MasterTx) error {
		projectSlice, err = i.projectService.GetAll(ctx, masterTx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return projectSlice, nil
}
