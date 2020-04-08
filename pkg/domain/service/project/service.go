package project

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository/project"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateNewProject(ctx *gin.Context, masterTx mysql.MasterTx, userID, title, description string) error
	GetByPK(ctx *gin.Context, masterTx mysql.MasterTx, id int) (*model.Project, error)
	GetByUserID(ctx *gin.Context, masterTx mysql.MasterTx, userID string) (model.ProjectSlice, error)
	GetAll(ctx *gin.Context, masterTx mysql.MasterTx) (model.ProjectSlice, error)
}

type service struct {
	projectRepository project.Repository
}

func New(projectRepository project.Repository) Service {
	return &service{
		projectRepository: projectRepository,
	}
}

func (s *service) CreateNewProject(ctx *gin.Context, masterTx mysql.MasterTx, userID, title, description string) error {
	if err := s.projectRepository.InsertProject(ctx, masterTx, &model.Project{
		UserID:      userID,
		Title:       title,
		Description: description,
	}); err != nil {
		return err
	}
	return nil
}

func (s *service) GetByPK(ctx *gin.Context, masterTx mysql.MasterTx, id int) (*model.Project, error) {
	project, err := s.projectRepository.SelectByPK(ctx, masterTx, id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (s *service) GetByUserID(ctx *gin.Context, masterTx mysql.MasterTx, userID string) (model.ProjectSlice, error) {
	projects, err := s.projectRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (s *service) GetAll(ctx *gin.Context, masterTx mysql.MasterTx) (model.ProjectSlice, error) {
	return s.projectRepository.SelectAll(ctx, masterTx)
}
