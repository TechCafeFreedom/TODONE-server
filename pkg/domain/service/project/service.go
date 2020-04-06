package project

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository/project"
)

type Service interface {
	CreateNewProject(userID, title, description string) error
	GetByPK(id int) (*model.Project, error)
	GetByUserID(userID string) (model.ProjectSlice, error)
	GetAll() (model.ProjectSlice, error)
}

type service struct {
	projectRepository project.Repository
}

func New(projectRepository project.Repository) Service {
	return &service{
		projectRepository: projectRepository,
	}
}

func (s *service) CreateNewProject(userID, title, description string) error {
	if err := s.projectRepository.InsertProject(&model.Project{
		UserID:      userID,
		Title:       title,
		Description: description,
	}); err != nil {
		return err
	}
	return nil
}

func (s *service) GetByPK(id int) (*model.Project, error) {
	project, err := s.projectRepository.SelectByPK(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *service) GetByUserID(userID string) (model.ProjectSlice, error) {
	projects, err := s.projectRepository.SelectByUserID(userID)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *service) GetAll() (model.ProjectSlice, error) {
	return s.projectRepository.SelectAll()
}
