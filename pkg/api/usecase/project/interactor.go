package project

import (
	"errors"
	"todone/db/mysql/model"
	projectservice "todone/pkg/domain/service/project"

	"github.com/gin-gonic/gin"
)

type Interactor interface {
	CreateNewProject(context *gin.Context, title, description string) error
	GetByPK(id int) (*model.Project, error)
	GetByUserID(ctx *gin.Context) (model.ProjectSlice, error)
	GetAll() (model.ProjectSlice, error)
}

type intereractor struct {
	projectService projectservice.Service
}

func New(projectService projectservice.Service) Interactor {
	return &intereractor{
		projectService: projectService,
	}
}

func (i *intereractor) CreateNewProject(c *gin.Context, title, description string) error {
	userID, ok := c.Get("AUTHED_USER_ID")
	if !ok {
		return errors.New("userID is not found in context")
	}

	if err := i.projectService.CreateNewProject(userID.(string), title, description); err != nil {
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

func (i *intereractor) GetByUserID(c *gin.Context) (model.ProjectSlice, error) {
	userID, ok := c.Get("AUTHED_USER_ID")
	if !ok {
		return nil, errors.New("userID is not found in context")
	}

	projects, err := i.projectService.GetByUserID(userID.(string))
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
