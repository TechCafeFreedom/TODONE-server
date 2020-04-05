package project

import (
	model "todone/db/mysql/models"
	"todone/pkg/domain/repository/project"
)

type Service interface {
	CreateNewProject(userID, title, description string) error
	SelectAll() ([]*model.Project, error)
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

func (s *service) SelectAll() ([]*model.Project, error) {
	return s.projectRepository.SelectAll()
}
