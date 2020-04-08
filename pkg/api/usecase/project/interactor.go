package project

import (
	"todone/db/mysql/model"
	projectservice "todone/pkg/domain/service/project"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
)

type Interactor interface {
	CreateNewProject(ctx *gin.Context, userID, title, description string) error
	GetByPK(id int) (*model.Project, error)
	GetByUserID(userID string) (model.ProjectSlice, error)
	GetAll() (model.ProjectSlice, error)
}

type intereractor struct {
	masterTxManager mysql.MasterTxManager
	projectService  projectservice.Service
}

func New(masterTxManager mysql.MasterTxManager, projectService projectservice.Service) Interactor {
	return &intereractor{
		masterTxManager: masterTxManager,
		projectService:  projectService,
	}
}

func (i *intereractor) CreateNewProject(ctx *gin.Context, userID, title, description string) error {
	err := i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx mysql.MasterTx) error {
		if err := i.projectService.CreateNewProject(userID, title, description); err != nil {
			return err
		}
		if err := i.projectService.CreateNewProject(userID, title, description); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *intereractor) GetByPK(id int) (*model.Project, error) {
	project, err := i.projectService.GetByPK(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (i *intereractor) GetByUserID(userID string) (model.ProjectSlice, error) {
	projects, err := i.projectService.GetByUserID(userID)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (i *intereractor) GetAll() (model.ProjectSlice, error) {
	projects, err := i.projectService.GetAll()
	if err != nil {
		return nil, err
	}
	return projects, nil
}
