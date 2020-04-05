package project

import (
	"errors"
	model "todone/db/mysql/models"
	projectservice "todone/pkg/domain/service/project"

	"github.com/gin-gonic/gin"
)

type Interactor interface {
	CreateNewProject(context *gin.Context, title, description string) error
	SelectAll() ([]*model.Project, error)
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

func (i *intereractor) SelectAll() ([]*model.Project, error) {
	projects, err := i.projectService.SelectAll()
	if err != nil {
		return nil, err
	}
	return projects, nil
}
