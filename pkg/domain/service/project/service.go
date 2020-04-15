package project

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/project"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateNewProject(ctx *gin.Context, masterTx repository.MasterTx, userID, title, description string) error
	GetByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*model.Project, error)
	GetByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID string) (model.ProjectSlice, error)
	GetAll(ctx *gin.Context, masterTx repository.MasterTx) (model.ProjectSlice, error)
}

type service struct {
	projectRepository project.Repository
}

func New(projectRepository project.Repository) Service {
	return &service{
		projectRepository: projectRepository,
	}
}

func (s *service) CreateNewProject(ctx *gin.Context, masterTx repository.MasterTx, userID, title, description string) error {
	if err := s.projectRepository.InsertProject(ctx, masterTx, &model.Project{
		UserID:      userID,
		Title:       title,
		Description: description,
	}); err != nil {
		return err
	}
	return nil
}

func (s *service) GetByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*model.Project, error) {
	projectData, err := s.projectRepository.SelectByPK(ctx, masterTx, id)
	if err != nil {
		return nil, err
	}
	return projectData, nil
}

func (s *service) GetByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID string) (model.ProjectSlice, error) {
	projects, err := s.projectRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *service) GetAll(ctx *gin.Context, masterTx repository.MasterTx) (model.ProjectSlice, error) {
	return s.projectRepository.SelectAll(ctx, masterTx)
}
