package project

import (
	"todone/db/mysql/model"
	projectservice "todone/pkg/domain/service/project"
)

type Interactor interface {
	CreateNewProject(userID, title, description string) error
	GetByPK(id int) (*model.Project, error)
	GetByUserID(userID string) (model.ProjectSlice, error)
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

func (i *intereractor) CreateNewProject(userID, title, description string) error {
	if err := i.projectService.CreateNewProject(userID, title, description); err != nil {
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
