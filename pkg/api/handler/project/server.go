package project

import (
	"net/http"
	"todone/pkg/api/request/reqbody"
	"todone/pkg/api/response"
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

	res := response.ProjectResponse{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}

	context.JSON(http.StatusOK, res)
}

func (s *Server) GetProjectsByUserID(context *gin.Context) {
	projects, err := s.projectInteractor.GetByUserID(context)
	if err != nil {
		context.Error(err)
	}

	res := make(response.ProjectsResponse, 0, len(projects))
	for _, data := range projects {
		res = append(res, &response.ProjectResponse{
			ID:          data.ID,
			Title:       data.Title,
			Description: data.Description,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		})
	}

	context.JSON(http.StatusOK, res)
}

func (s *Server) GetAllProjects(context *gin.Context) {
	projects, err := s.projectInteractor.GetAll()
	if err != nil {
		context.Error(err)
	}

	res := make(response.ProjectsResponse, 0, len(projects))
	for _, data := range projects {
		res = append(res, &response.ProjectResponse{
			ID:          data.ID,
			Title:       data.Title,
			Description: data.Description,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		})
	}

	context.JSON(http.StatusOK, res)
}
