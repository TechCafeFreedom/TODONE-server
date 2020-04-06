package project

import (
	"errors"
	"net/http"
	"todone/db/mysql/model"
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

	userID, ok := context.Get("AUTHED_USER_ID")
	if !ok {
		context.Error(errors.New("userID is not found in context"))
	}

	if err := s.projectInteractor.CreateNewProject(userID.(string), reqBody.Title, reqBody.Description); err != nil {
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

	context.JSON(http.StatusOK, convertToProjectResponse(project))
}

func (s *Server) GetProjectsByUserID(context *gin.Context) {
	userID, ok := context.Get("AUTHED_USER_ID")
	if !ok {
		context.Error(errors.New("userID is not found in context"))
	}

	projects, err := s.projectInteractor.GetByUserID(userID.(string))
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, convertToProjectsResponse(projects))
}

func (s *Server) GetAllProjects(context *gin.Context) {
	projects, err := s.projectInteractor.GetAll()
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, convertToProjectsResponse(projects))
}

func convertToProjectResponse(project *model.Project) response.ProjectResponse {
	return response.ProjectResponse{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func convertToProjectsResponse(projects model.ProjectSlice) response.ProjectsResponse {
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
	return res
}
