package project

import (
	"net/http"
	"todone/pkg/api/request/reqbody"
	projectinteractor "todone/pkg/api/usecase/project"

	"github.com/gin-gonic/gin"
)

type Server struct {
	projectInteractor projectinteractor.Interactor
}

func New(projectInteractor projectinteractor.Interactor) Server {
	return Server{projectInteractor: projectInteractor}
}

func (s *Server) CreateNewProject(context *gin.Context) {
	var reqBody reqbody.ProjectCreate
	if err := context.BindJSON(&reqBody); err != nil {
		context.Error(err)
	}

	if err := s.projectInteractor.CreateNewProject(context, reqBody.Title, reqBody.Description); err != nil {
		context.Error(err)
	}

	context.Status(http.StatusOK)
}

func (s *Server) GetProjectByPK(context *gin.Context) {
	var reqBody reqbody.ProjectPK
	if err := context.BindJSON(&reqBody); err != nil {
		context.Error(err)
	}

	project, err := s.projectInteractor.GetByPK(reqBody.ID)
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, project)
}

func (s *Server) GetProjectsByUserID(context *gin.Context) {
	projects, err := s.projectInteractor.GetByUserID(context)
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, projects)
}

func (s *Server) GetAllProjects(context *gin.Context) {
	projects, err := s.projectInteractor.GetAll()
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"projects": projects})
}
